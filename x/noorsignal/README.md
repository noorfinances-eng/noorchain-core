# x/noorsignal — PoSS (Proof of Signal Social) Module (Placeholder)

This directory is reserved for the custom NOORCHAIN PoSS module.

The `noorsignal` module will be responsible for:

- Accepting and validating **social signals**:
  - donations
  - verified participation
  - certified content
  - other NOORCHAIN-specific actions

- Applying **anti-abuse rules**:
  - daily limits
  - curator checks
  - signal weighting (0.5x to 5x, etc.)

- Minting NUR according to:
  - fixed supply (299,792,458 NUR)
  - time-based halving every 8 years
  - automatic **70% / 30%** reward split:
    - 70% to the participant
    - 30% to the curator (final name TBD)

At this stage, this folder only holds documentation.
The actual implementation (protobuf, keeper, messages, queries)
will be added in later phases.
