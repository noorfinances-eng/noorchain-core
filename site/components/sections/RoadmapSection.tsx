export default function RoadmapSection() {
  return (
    <section
      id="roadmap"
      className="relative w-full bg-transparent scroll-mt-[84px]"
    >
      <div className="relative z-10 container pb-10 md:pb-14">
        <div className="max-w-3xl">
          {/* CONTENU INCHANGÉ */}
          
          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white/10 border border-white/25 backdrop-blur-md px-3 py-1 mb-6">
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
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/12 backdrop-blur-sm py-3 rounded-r-lg">
            The NOORCHAIN roadmap reflects executed milestones and planned work.
            Items are presented with explicit status labels and may evolve based
            on engineering validation and governance decisions.
          </p>

          {/* COMPLETED */}
          <div className="mb-6 p-5 border border-white/20 bg-white/10 backdrop-blur-md rounded-xl shadow-sm">
            <h2 className="text-xl font-semibold text-navy mb-2">✅ Completed</h2>
            <ul className="list-disc pl-6 text-gray-700 space-y-2">
              <li>M10 — Mainnet-like multi-node pack</li>
              <li>M11 — dApps v0 tooling</li>
              <li>M12 — World State + RPC Parity Foundations</li>
              <li>M13 — Ethereum-compatible execution hardening (receipts/logs persistence, state-by-block reads)</li>
              <li>M14 — LogRec index + eth_getLogs (range-based, ordering stability, compat/perf gates)</li>
              <li>PHASE 7 — Proof-of-liveness (frozen)</li>        
            </ul>
          </div>

          {/* IN PROGRESS */}
          <div className="mb-6 p-5 border border-white/20 bg-white/10 backdrop-blur-md rounded-xl shadow-sm">
            <h2 className="text-xl font-semibold text-navy mb-2">🔧 In progress</h2>
            <ul className="list-disc pl-6 text-gray-700 space-y-2">
              <li>M15 — JSON-RPC Filters (eth_newFilter / eth_getFilterChanges / eth_getFilterLogs / eth_uninstallFilter + block filter)</li>
              <li>RPC compatibility hardening (wallet/dApp edge-cases, error model consistency)</li>
              <li>Operational stability (restart invariants, leader/follower parity under load)</li>
            </ul>
          </div>

          {/* PLANNED */}
          <div className="p-5 border border-white/20 bg-white/10 backdrop-blur-md rounded-xl shadow-sm">
            <h2 className="text-xl font-semibold text-navy mb-2">⏳ Planned</h2>
            <ul className="list-disc pl-6 text-gray-700 space-y-2">
              <li>Extended RPC coverage (subscriptions, additional wallet compatibility endpoints)</li>
              <li>Public testnet opening (gated; staged rollout)</li>
              <li>dApps ecosystem expansion</li>
              <li>Security review and audit readiness</li>
              <li>Partnerships and integrations</li>
            </ul>
          </div>

        </div>
      </div>
    </section>
  );
}
