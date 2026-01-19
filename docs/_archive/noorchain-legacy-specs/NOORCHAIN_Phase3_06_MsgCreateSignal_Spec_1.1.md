NOORCHAIN 1.0 — Phase 3.06
PoSS Message Specification — MsgCreateSignal
Version 1.1 — Final Interface
Last Updated: 2025-12-03
1. Purpose of the Message

MsgCreateSignal is the unique transaction type used to create a PoSS signal on-chain.

It represents a verified positive social action, produced by:

a participant (noor1… address)

validated by a curator (noor1… address)

belonging to one of the four official PoSS categories

optionally referencing external metadata (QR, hash, link…)

It is the core input to NOORCHAIN’s social mining engine.

MsgCreateSignal is the only entry point to the PoSS reward pipeline.

2. Message Structure (Final Interface)
MsgCreateSignal {
    participant   string   // noor1...
    curator       string   // noor1...
    signal_type   uint32   // enum: MICRO_DONATION / PARTICIPATION / CONTENT / CCN
    reference     string   // optional metadata, max 256 chars
}

Structural rules

participant != curator

both must decode as valid NOOR bech32 addresses

signal_type must belong to the official enum

reference is optional but MUST NOT exceed 256 characters

This interface is considered final and will not change.

3. Validation Rules

Validation occurs before any state transition.

3.1. Basic checks

participant is a valid sdk.AccAddress

curator is a valid sdk.AccAddress

signal type is valid

reference ≤ 256 characters

If any of these fail → transaction rejected.

3.2. Daily Limits

The keeper reads:

participant daily counter

curator daily counter

If either exceeds the limits defined in Params:

MaxSignalsPerDay

MaxSignalsPerCuratorPerDay

→ The message still succeeds,
→ but rewards = 0 for the affected party.

Counters still increment.
The message is not rejected.

3.3. PoSS Enabled

If:

PoSSEnabled == false


Then:

reward = 0

counters increment

events emitted normally

This allows PoSS to operate in observation mode during testnet.

4. State Transitions (To Be Implemented in Phase 4)

When MsgCreateSignal is fully wired, it will perform:

4.1. Increment Daily Counters

participantDailyCount +1

curatorDailyCount +1

4.2. Increment Global Counters

TotalSignals++

4.3. Compute Reward

Using:

BaseReward

weights

halving era

70/30 structural split

daily caps

PoSSEnabled state

4.4. Mint / Transfer NUR

(Phase 4/6 implementation only)

70% to participant

30% to curator

OR zero rewards if limits exceeded or PoSS disabled

4.5. Update TotalMinted
4.6. Emit PoSS Events

This section is documentation only — implementation occurs during Phase 4.

5. Events Produced

MsgCreateSignal emits a transparent event for indexers and explorers.

Example structure:

EventCreateSignal {
    participant
    curator
    signal_type
    reward_participant
    reward_curator
    halving_era
    reference
}


Visible to:

explorers

analytics dashboards

governance monitors

PoSS visualisation tools

6. Errors & Non-Errors
6.1. Errors (transaction rejected)

invalid participant address

invalid curator address

invalid signal type

reference too long

6.2. NOT Errors (transaction succeeds)

exceeding participant daily limit → reward = 0

exceeding curator daily limit → curator reward = 0

PoSSEnabled = false → reward = 0

This preserves:

real user activity

chain liveness

transparent accounting for PoSS statistics

7. Policy Decision (Final): Accept Without Reward

NOORCHAIN never rejects PoSS signals for economic limits.

It ALWAYS:

accepts the transaction

increments counters

sets reward = 0 if needed

This maintains:

fairness

transparency

no accidental chain halts

clean accounting for the PoSS 80% reserve

Legal Light CH compliance

8. Security & Anti-Abuse Invariants

The implementation MUST enforce:

8.1. Global Cap Always Preserved

PoSS MUST NOT mint above the fixed supply:
299,792,458 NUR.

8.2. Structural Reward Split

Must always remain:

70% participant

30% curator

Not modifiable by governance.

8.3. Reserve and Denom

PoSSReserveDenom = "unur"

no alternative denominations allowed

8.4. Halving Is Mandatory

Rewards MUST respect the 8-year halving schedule.

8.5. Daily Counters

cannot overflow

must not reset incorrectly

must always be consistent

8.6. No Bypass of Validation

All signals MUST pass through this message.

9. Implementation Notes (Phase 4)

MsgCreateSignal requires:

msg_server.go

complete keeper integration

state transitions

reward computation logic

BeginBlock/EndBlock integration

event emission

ante handler compatibility

inclusion in module routing

This will be implemented in Phase 4B–4C.

End of PoSS Logic 16 — Message Specification Completed

This file is now final and ready for development in Phase 4.
