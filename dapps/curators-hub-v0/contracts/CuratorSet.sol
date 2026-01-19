// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title CuratorSet
/// @notice Minimal curator registry + threshold for PoSS v0 (application layer, not consensus).
contract CuratorSet {
    mapping(address => bool) private _isCurator;
    uint256 private _curatorCount;
    uint256 private _threshold;
    address public admin;

    event CuratorAdded(address indexed curator);
    event CuratorRemoved(address indexed curator);
    event ThresholdUpdated(uint256 threshold);
    event AdminUpdated(address indexed admin);

    modifier onlyAdmin() {
        require(msg.sender == admin, "CuratorSet: not admin");
        _;
    }

    constructor(address initialAdmin, address[] memory initialCurators, uint256 initialThreshold) {
        require(initialAdmin != address(0), "CuratorSet: admin=0");
        admin = initialAdmin;
        emit AdminUpdated(initialAdmin);

        for (uint256 i = 0; i < initialCurators.length; i++) {
            _addCurator(initialCurators[i]);
        }

        _setThreshold(initialThreshold);
    }

    function addCurator(address curator) external onlyAdmin {
        _addCurator(curator);
    }

    function removeCurator(address curator) external onlyAdmin {
        require(curator != address(0), "CuratorSet: curator=0");
        require(_isCurator[curator], "CuratorSet: not curator");
        _isCurator[curator] = false;
        _curatorCount -= 1;

        // Ensure threshold remains valid.
        if (_threshold > _curatorCount) {
            _threshold = _curatorCount;
            emit ThresholdUpdated(_threshold);
        }

        emit CuratorRemoved(curator);
    }

    function setThreshold(uint256 t) external onlyAdmin {
        _setThreshold(t);
    }

    function setAdmin(address newAdmin) external onlyAdmin {
        require(newAdmin != address(0), "CuratorSet: admin=0");
        admin = newAdmin;
        emit AdminUpdated(newAdmin);
    }

    function isCurator(address who) external view returns (bool) {
        return _isCurator[who];
    }

    function getThreshold() external view returns (uint256) {
        return _threshold;
    }

    function getCuratorCount() external view returns (uint256) {
        return _curatorCount;
    }

    function _addCurator(address curator) internal {
        require(curator != address(0), "CuratorSet: curator=0");
        require(!_isCurator[curator], "CuratorSet: already curator");
        _isCurator[curator] = true;
        _curatorCount += 1;
        emit CuratorAdded(curator);

        // If this is the first curator and threshold was zero, keep it safe for later explicit set.
        if (_threshold == 0 && _curatorCount == 1) {
            // no-op; threshold must be set explicitly via constructor/setThreshold
        }
    }

    function _setThreshold(uint256 t) internal {
        require(_curatorCount > 0, "CuratorSet: no curators");
        require(t >= 1 && t <= _curatorCount, "CuratorSet: bad threshold");
        _threshold = t;
        emit ThresholdUpdated(t);
    }
}
