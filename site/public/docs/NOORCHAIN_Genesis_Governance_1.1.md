NOORCHAIN_Genesis_Governance_1.1.md
NOORCHAIN FOUNDATION — Genesis Governance
Version 1.1
Last Updated: 10.12.2025

1. Purpose of Genesis Governance

This document defines the initial governance configuration embedded in the
Mainnet Genesis Block (genesis.json).

It ensures:

transparent initial allocations

clearly defined governance responsibilities

alignment with the Swiss Legal Light CH framework

mission-oriented distribution of early resources

correct initialization of PoSS parameters

proper mapping of institutional addresses

immutability of core governance principles

Genesis Governance forms the backbone of NOORCHAIN’s legitimacy, safety, and long-term stability.

In case of conflict, the Foundation Statutes and Legal Architecture prevail.

2. Immutable Genesis Principles

The following principles are permanently embedded and cannot be altered post-launch:

1. Fixed Supply

Total cap = 299,792,458 NUR (symbolic “speed of light”).

2. Halving Every 8 Years

PoSS issuance decreases according to protocol rules.
This is a technical rule, not a financial forecast.

3. No Inflation Beyond PoSS Reserve

No discretionary minting is permitted.

4. 5 / 5 / 5 / 5 / 80 Economic Allocation

5% Foundation

5% Noor Dev Sàrl

5% PoSS Stimulus

5% Pre-sale (optional, vested, private only)

80% PoSS Mintable Reserve

5. PoSS Reward Split (Immutable)

70% participant

30% curator

6. No Custody

Genesis governance cannot introduce custodial functions.

7. Non-Investment Token

Governance cannot transform NUR into a financial or yield-bearing instrument.

8. Legal Light CH Compatibility

All rules align with Swiss non-profit, non-custodial, non-speculative standards.

3. Genesis Addresses (To Be Finalized in Phase 7)

These placeholder addresses will be replaced with verified Bech32 addresses in Phase 7.

Genesis addresses are institutional assignment points.
They cannot receive user funds and do not function as financial accounts.

3.1 Foundation Address

noor1foundationxxxxxxxxxxxxxxxxxxxxx
Purpose: Holds 5% supply, controlled exclusively via 3/5 multisig.

3.2 Noor Dev Sàrl Address

noor1devsarlxxxxxxxxxxxxxxxxxxxxxxxxxxx
Purpose: Functional 5% allocation for R&D and infrastructure.

3.3 PoSS Stimulus Reserve

noor1stimulusxxxxxxxxxxxxxxxxxxxxxxxx
Purpose: Early ecosystem onboarding.

3.4 Pre-sale Optional Pool

noor1presalexxxxxxxxxxxxxxxxxxxxxxxxxx
Purpose: Optional pre-mainnet private allocation.

Clarification:
This does not represent a public offering, ICO, or token sale.
Any funds raised go exclusively to Noor Dev Sàrl, not the Foundation.

3.5 PoSS Mintable Reserve (80%)

noor1possreservexxxxxxxxxxxxxxxxxxxxxx
Purpose: Protocol-controlled PoSS issuance only.

4. Genesis Allocations (Hard-Coded)
Category	Percentage	Supply (NUR)
Foundation	5%	14,989,623
Noor Dev Sàrl	5%	14,989,623
PoSS Stimulus	5%	14,989,623
Pre-sale Optional	5%	14,989,623
PoSS Mintable Reserve	80%	239,833,984
Total	100%	299,792,458

These values are embedded in genesis.json.

5. Genesis Governance Powers
5.1 Foundation Board

Approves administrative actions

Oversees documentation & mission alignment

Supervises Curators

Publishes annual transparency reports

5.2 Multi-sig Committee (3/5)

Executes on-chain Foundation actions

Controls the Foundation allocation

Adjusts allowed PoSS parameters (never structural ones)

May temporarily disable PoSS in emergencies,
provided the action is documented, non-financial, and compliant with Legal Light CH

Implements governance-approved module deployments

5.3 Noor Dev Sàrl

No inherent governance authority

Holds functional allocation (5%)

Produces technical proposals

May submit governance proposals

5.4 Curators

No direct governance authority at genesis

Provide advisory input

Form the basis of future social governance models

6. Genesis Parameters (PoSS & Core Protocol)
6.1 Initial PoSS Parameters

Governance may adjust these post-launch except for immutable invariants:

PoSSEnabled: false

BaseReward: 1 unur (non-financial unit)

WeightMicroDonation: 2

WeightParticipation: 1

WeightContent: 3

WeightCCN: 4

MaxSignalsPerDay: 10

MaxSignalsPerCuratorPerDay: 20

MaxRewardPerDay: 100 unur (non-financial cap)

HalvingPeriodYears: 8

6.2 Economic Constraints

Minting cannot exceed the PoSS reserve

No external minting

No discretionary pool creation

Governance cannot modify: halving, supply, reward split, PoSS reserve size.

7. Governance Limits (Hard-Coded)

Governance cannot:

modify total supply

change PoSS structural split (70/30)

introduce inflation

alter genesis allocation percentages

circumvent Legal Light CH restrictions

reassign genesis addresses

mint tokens outside PoSS logic

8. Post-Genesis Upgrade Path

Possible future governance extensions:

Activation of PoSS after stability period

Addition of Curators

Adjustment of PoSS limits and weights

Deployment of ecosystem modules

Extended delegation or social governance mechanisms

Multisig rotation

Curator advisory participation

All must comply with Legal Light CH.

9. Documentation Requirements

Before or at mainnet launch, the following must be published:

Governance Charter (C2)

Multisig Charter (C1)

Legal Light PDF

Genesis Allocation PDF

Foundation Statutes

Together they form the Genesis Governance Pack.

10. Adoption

This document is adopted by the
NOORCHAIN Foundation Board
and included in the Genesis Pack for mainnet launch.

Signatures:

Version 1.1 — Governance Phase
