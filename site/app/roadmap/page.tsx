export default function RoadmapPage() {
  return (
    <main className="w-full bg-paper">
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">

          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white border border-gray-soft px-3 py-1 mb-6">
            <span className="h-2 w-2 rounded-full bg-primary" />
            <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
              Project Timeline
            </span>
          </div>

          {/* TITLE */}
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            Roadmap
          </h1>

          {/* INTRO */}
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/60 py-3 rounded-r-lg">
            The NOORCHAIN roadmap reflects executed milestones and planned work.
            Items are presented with explicit status labels and may evolve based
            on engineering validation and governance decisions.
          </p>

          {/* ROADMAP (STATUS-DRIVEN) */}
          <section className="space-y-6">
            {/* ✅ COMPLETED */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
              <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-2">
                <h2 className="text-xl font-semibold text-navy">
                  ✅ Completed
                </h2>
                <p className="text-sm text-gray-600">
                  Verified milestones (tagged builds)
                </p>
              </div>

              <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
                <li>
                  <span className="font-semibold text-navy">M10 — Mainnet-like multi-node pack</span>{" "}
                  (leader/follower, P2P, health, runbook) —{" "}
                  <span className="font-mono text-sm">M10-MAINNETLIKE-STABLE</span>
                </li>
                <li>
                  <span className="font-semibold text-navy">M11 — dApps v0 tooling and wallet compatibility fixes</span>{" "}
                  — <span className="font-mono text-sm">M11-DAPPS-STABLE</span>
                </li>
                <li>
                  <span className="font-semibold text-navy">M12.2 — World state groundwork</span>{" "}
                  (StateDB-backed reads: nonce/balance) —{" "}
                  <span className="font-mono text-sm">M12.2-WORLDSTATE-RPC-NONCE-BALANCE</span>
                </li>
                <li>
                  <span className="font-semibold text-navy">PHASE 7 — Proof-of-liveness baseline</span>{" "}
                  (single public endpoint <span className="font-mono text-sm">/liveness.json</span>; controlled operation) —{" "}
                  <span className="font-mono text-sm">Frozen</span>
                </li>
              </ul>
            </div>

            {/* 🔧 IN PROGRESS */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
              <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-2">
                <h2 className="text-xl font-semibold text-navy">
                  🔧 In progress
                </h2>
                <p className="text-sm text-gray-600">
                  Active engineering work (scope-controlled)
                </p>
              </div>

              <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
                <li>
                  <span className="font-semibold text-navy">Persistent Ethereum-compatible world state</span>{" "}
                  (code/storage/stateRoot per block) and RPC reads beyond nonce/balance
                </li>
                <li>
                  <span className="font-semibold text-navy">EVM execution hardening</span>{" "}
                  (real transaction effects, predictable state transitions, restart consistency)
                </li>
                <li>
                  <span className="font-semibold text-navy">Operational stability</span>{" "}
                  (controlled exposure, monitoring, and runbooks refinement)
                </li>
              </ul>
            </div>

            {/* ⏳ PLANNED */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
              <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-2">
                <h2 className="text-xl font-semibold text-navy">
                  ⏳ Planned
                </h2>
                <p className="text-sm text-gray-600">
                  Scheduled work after validation gates
                </p>
              </div>

              <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
                <li>
                  <span className="font-semibold text-navy">Public testnet opening</span>{" "}
                  after feature completeness and security review
                </li>
                <li>
                  <span className="font-semibold text-navy">Broader dApps and ecosystem expansion</span>{" "}
                  (progressive rollout)
                </li>
                <li>
                  <span className="font-semibold text-navy">External audits and partnerships</span>{" "}
                  aligned with legal and operational readiness
                </li>
                <li>
                  <span className="font-semibold text-navy">Interoperability and liquidity (optional)</span>{" "}
                  evaluated only after core stability
                </li>
              </ul>
            </div>
          </section>

          {/* END LINE */}
          <div className="mt-10 h-px w-full bg-gray-soft" />
        </div>
      </section>
    </main>
  );
}
