Tokenomics — Version 1.1 (Public Release)

English — Official Institutional Edition

BLOC 1 — Introduction & Fixed Supply Model

NOORCHAIN — Tokenomics Overview
Version 1.1 — Public Documentation
Last Updated: 08.12.2025

BLOC 1 — Introduction & Fixed Supply Model

1. Introduction

NOORCHAIN is a Swiss-built, socially oriented blockchain protocol designed around transparency, non-custodial architecture, and long-term sustainability.

The network operates on Cosmos SDK + Ethermint, integrating a unique social-layer mechanism called Proof of Signal Social (PoSS) — an ethical, non-financial rewards model grounded in real-world positive actions.

The NUR token is a utility token, not a financial product.

It is used for:
• submitting and validating PoSS actions
• interacting with dApps such as Curators Hub, NOOR Pay, and CCN Studio
• governance participation (non-financial)
• accessing ecosystem tools

NOORCHAIN follows the Swiss Legal Light CH framework, which requires:
• no custody of user assets
• no guaranteed returns
• no investment or speculative messaging
• full transparency in governance and documentation
• strict separation between Foundation, developers, and community actors

This Tokenomics document describes the economic architecture of the protocol in a public, formal, and regulator-friendly format.

In case of conflict, the Foundation Statutes and the NOORCHAIN Legal Architecture prevail.

2. Fixed Supply & Non-Inflation Principles
2.1 Total Supply (Immutable)

The total supply of NUR is permanently fixed at:

299,792,458 NUR

This quantity is a symbolic reference to the speed of light, reflecting the project’s commitment to transparency, universality, and physical invariants.

2.2 No Inflation — Ever

• No new tokens can be minted.
• Supply cannot be increased by governance or by any entity.
• No discretionary expansion of the PoSS reserve is permitted.
• The system operates under a strict non-inflationary regime.

2.3 Utility Token Classification

NUR is strictly classified as a utility token, not a payment token, security token, or investment instrument.

It is not designed for:
• yield generation
• savings or interest
• speculative financial use

Its primary role is to support protocol utility and access to services. It does not grant equity, dividends, claims on assets, or any form of financial participation rights.

3. Genesis Allocation Model (Immutable)

The entire supply of 299,792,458 NUR is allocated at genesis according to a permanent, Swiss-aligned distribution framework.

This structure, referred to as 5 / 5 / 5 / 5 / 80, is fixed and cannot be modified after network launch.

3.1 Allocation Overview
Allocation Category	Percentage	Amount (NUR)	Purpose
Foundation	5%	14,989,622.9	Governance & public stewardship
Development Entity (Sàrl)	5%	14,989,622.9	Technical development & infrastructure
PoSS Stimulus Pool	5%	14,989,622.9	Early ecosystem support
Pre-sale Reserve (optional)	5%	14,989,622.9	Regulated private fundraising only
PoSS Mintable Reserve	80%	239,833,966.4	Long-term PoSS issuance

These allocations are hard-coded in genesis.json and enforced by protocol rules.
They may not be pledged as collateral, used as loan security, or repurposed for financial engineering.

3.2 Allocation Purpose & Governance Rationale
Foundation (5%) — Non-Profit Stewardship

The Foundation is responsible for:
• governance and compliance
• public documentation
• NGO and curator partnerships
• audit funding and transparency reports

Funds may only support activities aligned with the Foundation’s non-profit mission.
They cannot be used for speculation, investments, or financial operations of any kind.

Development Entity (5%) — Functional Allocation

The Development Entity (a Swiss Sàrl) supports:
• protocol maintenance
• infrastructure operations
• ecosystem integrations
• long-term R&D

This allocation does not grant governance authority.
The Sàrl operates independently and may generate revenue through services, not through token manipulation or control over supply.

PoSS Stimulus Pool (5%) — Early Ecosystem Activation

This pool accelerates early adoption and onboarding:
• NGOs and schools
• verified curators
• pilot programs
• social initiatives aligned with PoSS

It is strictly non-profit, non-speculative, and managed under multi-sig supervision.

Pre-sale Reserve (5%) — Optional & Regulated

If activated, this reserve can only be used under:
• Swiss Legal Light compliance
• mandatory vesting schedules
• transparent governance conditions
• multi-sig approvals

Funds raised from this reserve are received exclusively by the Development Entity (Sàrl), not by the Foundation.

If no pre-sale occurs, the reserve may remain dormant or be addressed through governance procedures consistent with Legal Light CH (for example, permanently locked or burned), without altering total supply or core economic invariants.

The Pre-sale Reserve does not represent an ICO, a public token sale, or an investment product.

PoSS Mintable Reserve (80%) — Long-Term Issuance

This is the backbone of NOORCHAIN’s economic model.
It provides the tokens used by the PoSS mechanism over multiple emission eras.

Characteristics include:
• no direct transfers
• issuance only via PoSS rules
• halving cycles every 8 years
• transparent and predictable distribution
• complete prohibition of discretionary minting

This design aims to ensure long-term sustainability without inflation.
All PoSS rewards originate exclusively from this reserve.

3.3 Institutional Roles in the Economic Model

Foundation
Acts as the neutral, public-good governance entity.
Holds 5% allocation under multi-sig control.
Cannot operate custodial or regulated financial services.

Development Entity (Sàrl)
Holds 5% allocation.
Executes technical work and may generate revenue through services (SaaS, integrations).
Has no authority to change token supply or core economic rules.

Multi-sig Committee (3/5)
Executes all on-chain Foundation actions.
Cannot modify fixed supply, allocation percentages, PoSS structural rules, or Legal Light constraints.

Curators
Validate PoSS signals.
Receive 30% of each PoSS reward.
Hold no treasury and no direct governance authority.

Community & Participants
Engage with the network, generate PoSS signals, and use dApps.
Remain fully self-custodial at all times.

3.4 Immutable Governance Constraints

The following principles cannot be changed, even by governance:
• total supply (299,792,458 NUR)
• allocation percentages (5 / 5 / 5 / 5 / 80)
• PoSS reward split (70/30)
• halving period (8 years)
• prohibition of inflation
• prohibition of discretionary minting
• Legal Light CH compliance boundaries

These constraints ensure the protocol remains transparent, ethical, and resistant to manipulation.

BLOC 3 — PoSS Economic Model
4. Proof of Signal Social (PoSS) — Economic Model

PoSS (Proof of Signal Social) is NOORCHAIN’s socially oriented reward mechanism.
It distributes tokens based on verified positive actions within communities, schools, NGOs, and social networks.

PoSS is not a financial product, is not yield-bearing, and does not involve deposits, lock-ups, or APR/APY.
It is a protocol-defined social reward system, fully compatible with the Swiss Legal Light CH framework.

4.1 Immutable Reward Split (70/30)

Every validated PoSS action distributes rewards as follows:
• 70% to the participant submitting the action
• 30% to the curator validating the action

This split is hard-coded and cannot be changed by governance, upgrades, or parameter adjustments.
It ensures fairness, decentralization of social validation, and ethical participation incentives.

4.2 Reward Structure

PoSS rewards are based on a combination of:
• the BaseReward defined in genesis
• the weight of the action type
• the halving era
• the daily limits
• the size of the remaining PoSS Reserve

Reward formula (simplified conceptual form):

raw_reward = BaseReward × Weight(action_type)
halved_reward = raw_reward / (2^halving_era)
distributed_reward =
  70% × halved_reward to participant
  30% × halved_reward to curator

This ensures transparent, predictable, and capped token issuance.

PoSS rewards must not be interpreted as income, salary, or guaranteed financial return.

4.3 Action Type Weights (Genesis Values)
Action Type	Weight
Micro-donations	2
Participation	1
Content creation	3
Certified CCN	4

These weights reflect the relative social impact of each signal type.
They can be adjusted through governance (except for the 70/30 split).

4.4 Halving Every 8 Years

PoSS uses an 8-year mechanical halving schedule inspired by scientific cycles rather than financial speculation.

• Year 0 → full reward
• Year 8 → rewards divided by 2
• Year 16 → rewards divided by 4
• Year 24 → rewards divided by 8
• … and so on

The halving mechanism:
• reduces issuance over time
• extends the lifetime of the PoSS reserve
• provides predictable and transparent token distribution
• remains independent of market price or external financial performance

4.5 Daily Limits

To ensure fairness and prevent farming, PoSS defines:
• maximum number of signals per participant per day
• maximum number of validated signals per curator per day
• maximum rewards per participant per day
• maximum rewards per curator per day

These parameters are governance-adjustable within strict boundaries but never affect:
• the fixed supply
• the PoSS 70/30 split
• the halving cycle
• the Legal Light constraints

4.6 Non-Financial Positioning

PoSS is deliberately designed to avoid financial or investment characteristics.
It is:
• symbolic
• capped
• transparent
• non-performance-based
• not linked to deposits or yields
• not marketed as remuneration

Users do not stake tokens, do not lock balances, and do not receive interest.
Rewards originate only from the PoSS reserve and are tied exclusively to social impact actions.

This ensures strict compliance with Swiss regulatory expectations for non-financial utility tokens.

4.7 Transparency of Distribution

All PoSS-related distributions are:
• recorded on-chain
• traceable
• capped by the halving model
• limited by the PoSS reserve
• visible to the community

There are no hidden emissions, no discretionary minting, and no privileged actors.

5. Long-Term Emission Timeline

The PoSS Mintable Reserve — representing 80% of the total supply — is structured to support the protocol for multiple emission eras.

Its emission schedule follows a predictable, mathematically defined pattern based on 8-year halving periods.

This approach ensures:
• long-term sustainability
• transparent issuance
• gradual reduction of rewards
• alignment with ethical and non-speculative economic principles

5.1 Emission Eras

The PoSS distribution model divides the lifetime of the reserve into distinct eras:

Era	Years	Relative Emission	Description
Era 1	0–8	100% of base rewards	Initial activation period
Era 2	9–16	50% of initial	First halving reduces issuance
Era 3	17–24	25% of initial	Second halving period
Era 4	25–32	12.5% of initial	Third halving period
Era 5+	33+	Continuing halving	Long-tail emissions

This schedule is designed as a technical emission model to extend the PoSS reserve over multiple decades, depending on network adoption and actual use of daily limits. It does not represent a financial projection or promise of duration.

5.2 No Inflation & Predictability

NOORCHAIN’s emission architecture is deliberately simple and transparent:
• There is no inflation.
• All tokens used for PoSS originate from the single PoSS reserve established at genesis.
• The halving schedule is independent of demand, price, or market behavior.

The system avoids incentives resembling interest, growth promises, or yield, and thus remains compatible with Swiss non-financial classification.

5.3 Reserve Sustainability

The halving model ensures:
• controlled early distribution when adoption is growing
• progressively lower distribution as the ecosystem matures
• a long-tail emission pattern suitable for multi-era operation

This approach is inspired by energy-constrained and resource-constrained systems rather than by financial markets, emphasizing:
• scarcity
• predictability
• ethical use
• sustainability

5.4 Alignment With Legal Light CH

The emission timeline has been structured to avoid:
• speculative cycles
• market-influenced issuance
• yield-based mechanics
• any perception of “expected returns”

Instead, the schedule supports:
• social participation
• community empowerment
• non-profit mission alignment
• transparent governance

5.5 Transition to Maturity

As PoSS emissions decline across halving periods, network utility is expected to shift toward:
• dApps usage
• community-driven value
• partnerships
• organic ecosystem growth

The protocol does not rely on constant token issuance to incentivize participation, which supports long-term economic and regulatory stability.

6. Governance and Economic Safety Architecture

The NOORCHAIN economic model is reinforced by a governance framework designed to protect the protocol, preserve Legal Light CH compliance, and guarantee long-term integrity.

This framework establishes strict boundaries, ensuring that neither governance bodies nor external actors can modify core economic rules.

6.1 Immutable Governance Constraints

The following elements are permanently fixed and cannot be altered by any governance process, upgrade, multi-sig action, or voting mechanism:

Total Supply
The supply is capped forever at 299,792,458 NUR.

Genesis Allocation Percentages
The 5 / 5 / 5 / 5 / 80 distribution is locked.

PoSS Structural Split
The 70/30 distribution (participant/curator) is immutable.

Halving Cycle
The 8-year halving period cannot be changed.

No Discretionary Minting
All issuance must come strictly from the PoSS reserve under PoSS rules.

Legal Light Restrictions
Rules preventing custody, yield mechanisms, or financial products cannot be overridden.

These constraints ensure the protocol remains predictable, transparent, and legally compliant.

6.2 Governance Powers (Restricted & Monitored)

Governance can adjust certain operational parameters within strict boundaries, such as:
• PoSS daily limits
• action type weights
• base reward amount
• curator onboarding rules
• module upgrades
• dApp integrations
• validator parameters

However, governance cannot:
• increase token supply
• alter foundational economic invariants
• repurpose the PoSS reserve
• assign tokens outside of genesis-defined pools
• weaken Legal Light protections

This protects the protocol from economic manipulation and regulatory risk.

6.3 Legal Light CH: Economic Compliance Boundaries

To maintain Swiss non-financial classification, NOORCHAIN must avoid:
• investment-like messaging
• any form of yield-based product
• custody of user funds
• fiat conversions operated internally
• speculative economic structures
• APR/APY language
• revenue promises
• Foundation-operated liquidity pools
• market-making activities

These restrictions apply to:
• the Foundation
• the Sàrl
• the Multi-sig Committee
• ecosystem partners
• all official communications

They form a non-negotiable compliance perimeter.

6.4 Allowed Economic Activities

NOORCHAIN supports a modern utility-token ecosystem within compliant boundaries.
The following activities are permitted:

• Orderbook exchange listing (e.g., MEXC)
• Free-market trading facilitated externally
• Utility-driven demand through dApps
• Partnerships with NGOs, schools, and institutions
• External PSP conversions (e.g., Mt Pelerin, NOWPayments)
• Grants and innovation funding
• SaaS revenue generated by the Development Entity
• Public educational and social-impact programs

None of these activities guarantee financial return or involve custodial operations.

6.5 Explicitly Forbidden Economic Activities

To preserve the mission and regulatory clarity, the following actions are permanently prohibited:

• Loans, interest, or credit mechanisms
• Internal yield farming or staking return products
• Foundation-operated liquidity pools
• Token sales marketed as investments
• Guaranteed token performance
• Token issuance outside PoSS rules
• Custody of user assets
• Price influence or market manipulation
• Misleading marketing or speculative claims

These prohibitions apply to all governance bodies and all future versions of NOORCHAIN.

6.6 Foundation vs. Development Entity — Separation of Powers

The economic model relies on a strict structural separation:

Foundation (Non-Profit)
• manages its 5% pool
• oversees governance
• conducts audits, compliance, documentation
• supports public and social missions
• cannot conduct commercial or speculative activity

Development Entity (Sàrl)
• manages its 5% pool
• is responsible for protocol development and infrastructure
• may generate revenue independently (SaaS, partnerships)
• cannot access Foundation funds
• cannot influence token supply or core economic invariants

This dual-entity structure ensures transparency, neutrality, and professional technical stewardship.

6.7 Multi-sig Committee — Economic Safety Layer

The multi-sig (3/5) enforces on-chain governance decisions.

It can:
• execute approved upgrades
• adjust PoSS parameters within defined limits
• manage the Foundation address
• supervise early-stage operations

It cannot:
• modify supply
• override genesis rules
• conduct financial operations
• act without transparent documentation

The multi-sig is an execution body, not a decision-making authority.

7. Market Model and Exchange Principles

NOORCHAIN adopts a market approach that is intentionally conservative, compliant, and aligned with Swiss Legal Light CH.

The goal is to enable free-market accessibility without transforming the project into a financial product or speculative instrument.

7.1 Allowed Market Structure

The following elements are permitted and compatible with the protocol’s utility nature:

Orderbook Exchange Listing
Listings on centralized exchanges (e.g., MEXC) operating orderbook systems are acceptable. These venues:
• do not require liquidity pools operated by NOORCHAIN
• permit market access without financial commitments from the Foundation
• allow decentralized price discovery through independent actors

External Liquidity Provision
If liquidity exists, it must originate from:
• independent market actors
• fully external DEX or CEX systems
• private initiatives

Neither the Foundation nor the Development Entity may act as liquidity providers.

Utility-Driven Demand
Use of NUR within:
• PoSS participation
• Curators Hub
• NOOR Pay
• CCN Studio
• educational or social programs

creates natural demand linked to utility, not to financial engineering.

External PSP Conversions
Conversions between NUR ↔ fiat are only allowed through regulated providers (e.g., Mt Pelerin or NOWPayments).

NOORCHAIN:
• does not operate on/off-ramp systems
• does not custody fiat or crypto
• does not participate directly in KYC/AML processes

7.2 Forbidden Market Mechanisms

To maintain strict regulatory clarity, NOORCHAIN prohibits:

• internal liquidity pools
• staking APY/APR products
• guaranteed returns or financial performance
• Foundation or Sàrl-led market-making
• token buybacks
• investment-oriented communication
• leveraged or synthetic financial instruments

These actions would compromise Legal Light compliance and create unacceptable regulatory exposure.

8. Founder Compensation and Economic Conduct

NOORCHAIN incorporates a transparent, ethically grounded model for leadership participation, ensuring alignment with long-term project health and Swiss governance norms.

8.1 No Special Access to Funds

The founder or any leadership role:
• cannot access Foundation funds
• cannot modify supply
• cannot direct PoSS issuance
• cannot receive preferential allocation

All compensation must follow formal, transparent channels.

8.2 Development Entity Allocation (5%)

The only structural economic participation is through the Development Entity’s 5% allocation, which follows:
• vesting rules
• independent corporate governance
• clear separation from Foundation assets

This ensures operational sustainability without centralizing economic power.

8.3 Responsible Market Behavior

To preserve ecosystem trust and avoid any perception of market manipulation, governance may enforce:
• a recommended maximum 0.5% monthly sale rate for Sàrl-held tokens
• transparent disclosures for significant transfers
• multi-sig oversight of vested token movements where appropriate

These rules enhance economic stability and maintain institutional confidence.

8.4 No Influence on Token Price

Leadership and governance bodies must refrain from:
• price-related communication
• performance claims
• speculative commentary
• coordinated buying or selling

The project must remain mission-first, not profit-driven.

9. Economic Integrity & Sustainability

The NOORCHAIN economic model supports:
• multi-decade longevity through the PoSS halving schedule
• strict non-inflation to protect token scarcity
• ethical distribution focused on community and public-good impact
• clear separation of roles between governance and development
• transparent reporting and regulatory alignment
• traceable on-chain distribution for all PoSS rewards

The system avoids dependency on token issuance or speculative cycles.
Instead, it fosters sustainable growth through utility, education, and social value.

10. Closing Summary

NOORCHAIN’s tokenomics combine fixed supply, predictable emission, structural fairness, and strict compliance to produce a model that is:
• transparent
• predictable
• non-speculative
• utility-based
• human-centered
• governance-safe
• aligned with Swiss ethical digital asset principles

This architecture is designed to keep the protocol viable over multiple emission eras and support a robust ecosystem of dApps, curators, community participants, and partners.

11. Final Conclusion

NOORCHAIN’s tokenomics represent a principled, long-term approach to blockchain economics.

The model balances:
• fixed supply and scarcity
• ethical distribution
• non-financial utility
• legal compliance
• transparent governance
• multi-decade sustainability
• social impact orientation
• decentralized participation
• strict separation of powers
• predictable PoSS emission cycles

By combining a capped supply, immutable allocation rules, and a halving-based distribution schedule, the economic system avoids inflationary risk.

Its architecture prevents speculative behavior while enabling genuine utility through dApps, PoSS interactions, and ecosystem partnerships.

The result is a resilient, Swiss-aligned blockchain economy designed for public good, real-world participation, and transparent governance — not for yield, speculation, or financial engineering.

This document forms a foundational component of the NOORCHAIN public documentation suite and should be referenced alongside:
• Genesis Pack 1.1
• Governance & Legal Framework
• Compliance Framework
• Foundation Statutes
• PoSS Technical Specifications
• Whitepapers v1.1

Together, these provide a complete and coherent economic and governance blueprint for NOORCHAIN 1.0 and its multi-decade roadmap.

12. Document Metadata & Versioning

Document Title: Tokenomics — Public Overview
Version: 1.1 (Public Release)
Language: English
Status: Stable
Maintained by: NOORCHAIN Foundation — Governance & Documentation Division
Last Updated: 08.12.2025
Document Type: Public reference, non-binding for investment purposes
Location: /site/public/docs/Tokenomics_Public_1.1.md
License: Open documentation (non-commercial reuse permitted with attribution)

Revision Policy
Updates to this document may occur only under the following conditions:
• clarifications in language
• updates following governance-approved parameter adjustments
• formatting or structural improvements

Core economic principles (supply, allocation, PoSS structure, halving, Legal Light constraints) cannot be changed.
