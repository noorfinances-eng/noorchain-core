Proof of Signal Social (PoSS) — Technical Specification (Phase 3)

Version: 1.1

1. Concept Summary

PoSS is a social-consensus mechanism where real-world positive actions generate rewards.

Three categories of signals:

Micro-donations

Verified participation (QR, events, social actions)

Certified content (videos, posts, creations validated by Curators)

2. Minable Supply

Cap: 299,792,458 NUR

Halving period: every 8 years

Reward distribution:

70% → participant

30% → curator

Emission controlled by the x/noorsignal module

3. Types of Signals (v1)
Type	Weight	Description
Micro donation	1×	Small on-chain donation
Verified Attendance	2×	QR-based confirmation at events
Verified Action	3×	Real action validated by Curator
Certified Content	5×	Video/post validated by Curator
4. Constraints

Daily limits

Curator validation required

Anti-abuse rules

Halving schedule enforced automatically

On-chain emission (no inflation)

5. Storage Model (planned)

signals/{day}/{address}

curators/{id}

params/

reward_pool