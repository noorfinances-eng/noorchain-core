export default function GenesisPage() {
  return (
    <main className="w-full bg-paper">
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">

          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white border border-gray-soft px-3 py-1 mb-6">
            <span className="h-2 w-2 rounded-full bg-primary" />
            <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
              Genesis & Allocation
            </span>
          </div>

          {/* TITLE */}
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            Genesis Overview
          </h1>

          {/* INTRO */}
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/60 py-3 rounded-r-lg">
            The NOORCHAIN genesis defines the fixed supply, PoSS distribution
            rules, halving schedule, and the 5 / 5 / 5 / 5 / 80 allocation
            model that structures long-term sustainability.
          </p>

          {/* SUPPLY */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Total Supply
            </h2>
            <p className="text-gray-700 leading-relaxed">
              The total supply is fixed at{" "}
              <strong>299,792,458&nbsp;NUR</strong>. This supply can never be
              increased. All PoSS rewards and allocations are derived from this
              fixed cap, ensuring predictable and auditable issuance over time.
            </p>
          </section>

          {/* DISTRIBUTION */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Initial Distribution
            </h2>
            <p className="text-gray-700 leading-relaxed mb-3">
              At genesis, the supply is structured into five clearly defined
              components:
            </p>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <strong>5%</strong> — NOOR Foundation
              </li>
              <li>
                <strong>5%</strong> — Dev Pool
              </li>
              <li>
                <strong>5%</strong> — PoSS Stimulus
              </li>
              <li>
                <strong>5%</strong> — Optional Pre-sale
              </li>
              <li>
                <strong>80%</strong> — PoSS Mintable Supply
              </li>
            </ul>
          </section>

          {/* GOVERNANCE PRINCIPLES */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Governance Principles
            </h2>
            <p className="text-gray-700 leading-relaxed mb-3">
              Genesis governance establishes immutable foundations for NOORCHAIN:
            </p>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>Fixed total supply cannot be changed.</li>
              <li>Halving of PoSS issuance occurs every 8 years.</li>
              <li>
                PoSS rewards always follow the{" "}
                <strong>70% participant / 30% curator</strong> split.
              </li>
              <li>
                All parameters remain transparent, on-chain and publicly
                auditable.
              </li>
            </ul>
          </section>

          {/* DOWNLOAD / GENESIS PACK */}
          <section className="p-6 border border-gray-soft rounded-xl bg-white shadow-sm mb-10 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Genesis Pack
            </h2>
            <p className="text-gray-700 leading-relaxed mb-4">
              The full genesis documentation, including governance, allocation,
              economic model and institutional structure, will be available as a
              public Genesis Pack document.
            </p>
            <a
              href="#"
              className="inline-block px-6 py-3 border border-primary text-primary rounded-md text-sm md:text-base font-medium hover:bg-primary hover:text-white transition"
            >
              Download Genesis Pack
            </a>
          </section>

          {/* END LINE */}
          <div className="h-px w-full bg-gray-soft" />
        </div>
      </section>
    </main>
  );
}
