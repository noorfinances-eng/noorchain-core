This document defines the complete specification for the PoSS transaction message:

MsgCreateSignal


This message will be implemented later in Phase 4 but its interface and behaviour are final.

1. Purpose of the message

MsgCreateSignal is the only message used to create a PoSS signal on-chain.

It represents:

a social action by a participant (noor1…)

validated by a curator (noor1…)

of a specific signal type

which may produce PoSS rewards according to:

weights,

halving era,

daily limits,

PoSSEnabled.

This message is the backbone of NOORCHAIN’s social mining.

2. Message structure (final)
message MsgCreateSignal {
    string participant = 1;   // noor1...
    string curator     = 2;   // noor1...
    uint32 signal_type = 3;   // enum: MICRO_DONATION / PARTICIPATION / CONTENT / CCN
    string reference   = 4;   // optional metadata (hash, QR, link, ID...)
}

Additional rules

participant != curator

both addresses MUST be valid NOOR bech32

signal_type MUST be one of the official enum

reference MAY be empty but MUST NOT exceed 256 characters

3. Validation (before state changes)
3.1. Basic checks

addresses must decode successfully as sdk.AccAddress

signal type must be valid

reference valid (<= 256 chars)

3.2. Daily limits

read daily counter for participant:

if >= MaxSignalsPerDay → reward disabled

read daily counter for curator:

if >= MaxSignalsPerCuratorPerDay → curator reward disabled

The message itself is not rejected unless governance decides otherwise (see §7).

3.3. PoSS Enabled

if PoSSEnabled == false → rewards = 0, counters still increment.

4. State transitions (Phase 4 implementation)

When fully implemented, the message will:

Increment daily counters

participant daily counter +1

curator daily counter +1

Increment global counters

TotalSignals++

Compute reward

call ComputeSignalReward(params, signalType, blockHeight)

apply halving

apply daily caps

Mint / transfer NUR

participant gets 70 %

curator gets 30 %

unless limited / disabled

Update TotalMinted

Emit events

This section is documentation only — the code will be written in Phase 4 PoSS Logic.

5. Events emitted
type EventCreateSignal struct {
    participant       string
    curator           string
    signal_type       string
    reward_participant string // ex: "5unur"
    reward_curator     string
    halving_era        uint64
    reference          string
}


These events allow explorers, dApps and dashboards to track PoSS activity transparently.

6. Errors & return types

The message MAY return:

ErrInvalidAddress

ErrInvalidSignalType

ErrReferenceTooLong

It MUST NOT fail for:

exceeding daily limits
→ rewards = 0 but signal is still processed.

PoSS disabled
→ rewards = 0 but signal is processed.

7. Policy: Reject or Accept-without-reward?

NOORCHAIN chooses:

✔️ Accept the signal
✔️ Increment counters
✔️ Reward = 0 if above limits or PoSS disabled

NEVER reject the transaction, unless data is malformed.

This ensures:

real user behaviour,

anti-abuse protection without chain halting,

fair long-term accounting for the 80% PoSS mining reserve.

8. Security & anti-abuse invariants

To be respected during implementation:

PoSS never mints above the global cap.

ParticipantShare + CuratorShare MUST always equal 100.

PoSSReserveDenom must be "unur".

Halving schedule MUST be applied to the raw reward.

No double-spend of counters.

Curators may not validate infinite signals:

limited by MaxSignalsPerCuratorPerDay.

No signal may bypass validation through MsgCreateSignal.

9. Implementation notes (Phase 4)

This message will require:

msg_server.go

integration with keeper

state updates

event emission

ante check compatibility

integration into app’s module manager

All will be implemented cleanly using the specification above.

✔️ End of PoSS Logic 16

(Ready for Phase 4 implementation in upcoming steps)