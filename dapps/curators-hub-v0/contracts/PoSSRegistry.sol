// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./CuratorSet.sol";

/// @title PoSSRegistry
/// @notice Minimal snapshot registry for PoSS v0: stores snapshot hash + metadata and verifies curator threshold signatures.
contract PoSSRegistry {
    struct Snapshot {
        bytes32 snapshotHash; // keccak256(canonical_minified_manifest_json_bytes)
        string uri;           // IPFS CID or HTTPS URL
        uint64 periodStart;
        uint64 periodEnd;
        uint64 publishedAt;
        uint32 version;       // manifest version (v1 for PoSS v0)
        address publisher;
    }

    struct SnapshotMeta {
        bytes32 snapshotHash;
        string uri;
        uint64 periodStart;
        uint64 periodEnd;
        uint32 version;
    }

    struct Signature {
        uint8 v;
        bytes32 r;
        bytes32 s;
    }

    CuratorSet public curatorSet;
    uint256 public snapshotCount;

    mapping(uint256 => Snapshot) private _snapshots;
    mapping(bytes32 => bool) private _knownSnapshotHash;

    event SnapshotSubmitted(
        uint256 indexed id,
        bytes32 indexed snapshotHash,
        string uri,
        uint64 periodStart,
        uint64 periodEnd,
        uint64 publishedAt,
        uint32 version,
        address indexed publisher
    );

    constructor(address curatorSetAddress) {
        require(curatorSetAddress != address(0), "PoSSRegistry: curatorSet=0");
        curatorSet = CuratorSet(curatorSetAddress);
    }

    function submitSnapshot(SnapshotMeta calldata meta, Signature[] calldata sigs) external returns (uint256 id) {
        require(meta.snapshotHash != bytes32(0), "PoSSRegistry: hash=0");
        require(meta.periodStart < meta.periodEnd, "PoSSRegistry: bad period");
        require(!_knownSnapshotHash[meta.snapshotHash], "PoSSRegistry: duplicate");

        uint256 threshold = curatorSet.getThreshold();
        require(threshold > 0, "PoSSRegistry: threshold=0");
        require(sigs.length >= threshold, "PoSSRegistry: not enough sigs");

        // Message that curators sign (EIP-191 simple): keccak256(chainId || snapshotHash)
        bytes32 msgHash = _messageHash(block.chainid, meta.snapshotHash);

        // Verify unique curator signers meeting threshold.
        address[] memory seen = new address[](sigs.length);
        uint256 validCount = 0;

        for (uint256 i = 0; i < sigs.length; i++) {
            address signer = ecrecover(msgHash, sigs[i].v, sigs[i].r, sigs[i].s);
            if (signer == address(0)) continue;

            // Must be curator.
            if (!curatorSet.isCurator(signer)) continue;

            // Must be unique.
            bool dup = false;
            for (uint256 j = 0; j < validCount; j++) {
                if (seen[j] == signer) {
                    dup = true;
                    break;
                }
            }
            if (dup) continue;

            seen[validCount] = signer;
            validCount += 1;

            if (validCount >= threshold) break;
        }

        require(validCount >= threshold, "PoSSRegistry: threshold not met");

        snapshotCount += 1;
        id = snapshotCount;

        _knownSnapshotHash[meta.snapshotHash] = true;

        Snapshot memory s = Snapshot({
            snapshotHash: meta.snapshotHash,
            uri: meta.uri,
            periodStart: meta.periodStart,
            periodEnd: meta.periodEnd,
            publishedAt: uint64(block.timestamp),
            version: meta.version,
            publisher: msg.sender
        });

        _snapshots[id] = s;

        emit SnapshotSubmitted(
            id,
            meta.snapshotHash,
            meta.uri,
            meta.periodStart,
            meta.periodEnd,
            uint64(block.timestamp),
            meta.version,
            msg.sender
        );
    }

    function getSnapshot(uint256 id) external view returns (Snapshot memory) {
        require(id >= 1 && id <= snapshotCount, "PoSSRegistry: bad id");
        return _snapshots[id];
    }

    function latestSnapshotId() external view returns (uint256) {
        return snapshotCount;
    }

    function isKnown(bytes32 snapshotHash) external view returns (bool) {
        return _knownSnapshotHash[snapshotHash];
    }

    function _messageHash(uint256 chainId, bytes32 snapshotHash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(chainId, snapshotHash));
    }
}
