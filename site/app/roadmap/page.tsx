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
            The NOORCHAIN roadmap outlines the major development phases of the
            project. Timelines are indicative and depend on technical progress
            and governance decisions.
          </p>

          {/* PHASES LIST */}
          <section className="space-y-6">
            
            {/* PHASE 1 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 1 — Framing & Decisions
              </h2>
              <p className="text-gray-700">Status: Completed</p>
            </div>

            {/* PHASE 2 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 2 — Core Skeleton
              </h2>
              <p className="text-gray-700">Status: Completed</p>
            </div>

            {/* PHASE 3 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 3 — Documentation
              </h2>
              <p className="text-gray-700">Status: Completed</p>
            </div>

            {/* PHASE 4 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 4 — Implementation & PoSS Module
              </h2>
              <p className="text-gray-700">Status: Completed</p>
            </div>

            {/* PHASE 5 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 5 — Legal & Governance
              </h2>
              <p className="text-gray-700">Status: Completed</p>
            </div>

            {/* PHASE 6 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 6 — Genesis Pack & Communication
              </h2>
              <p className="text-gray-700">Status: In progress</p>
            </div>

            {/* PHASE 7 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 7 — Mainnet
              </h2>
              <p className="text-gray-700">Status: Planned</p>
            </div>

            {/* PHASE 8 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 8 — dApps & Ecosystem
              </h2>
              <p className="text-gray-700">Status: Planned</p>
            </div>

            {/* PHASE 9 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 9 — Partnerships & Audits
              </h2>
              <p className="text-gray-700">Status: Planned</p>
            </div>

            {/* PHASE 10 */}
            <div className="p-5 border border-gray-soft bg-white rounded-xl shadow-sm">
              <h2 className="text-xl font-semibold text-navy mb-1">
                Phase 10 — Interoperability & Liquidity (Optional)
              </h2>
              <p className="text-gray-700">Status: Planned</p>
            </div>

          </section>

          {/* END LINE */}
          <div className="mt-10 h-px w-full bg-gray-soft" />
        </div>
      </section>
    </main>
  );
}
