export default function TechnologySection() {
  return (
    <section
      id="technology"
      className="relative w-full bg-transparent scroll-mt-[84px]"
    >
      <div className="relative z-10 container pt-4 pb-10 md:pt-6 md:pb-14">
        <div className="max-w-3xl">
          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white/10 border border-white/25 backdrop-blur-md px-3 py-1 mb-6">
            <span className="h-2 w-2 rounded-full bg-primary" />
            <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
              Core Architecture
            </span>
          </div>

          {/* TITLE */}
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            Technology
          </h1>

          {/* INTRO */}
          <p className="text-lg text-gray-700 leading-relaxed mb-8 border-l-4 border-primary pl-4 bg-white/12 backdrop-blur-sm py-3 rounded-r-lg">
            NOORCHAIN is built using a modular blockchain architecture designed for
            security, transparency, and verifiable social value.
          </p>

          {/* CORE STACK */}
          <section className="mb-6 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Core Stack
            </h2>
            <p className="text-gray-700 leading-relaxed">
              The network combines a core ledger, an EVM-compatible execution layer,
              and a dedicated PoSS module. This structure ensures transparent
              participation rules while allowing developers to build applications
              on top of a robust foundation.
            </p>
          </section>

          {/* POSS INTEGRATION */}
          <section className="mb-6 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              PoSS Integration
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Proof of Signal Social is natively embedded into the protocol. It applies
              daily limits, computes rewards, and validates participant–curator
              interactions according to public parameters, focusing on participation
              and curatorship rather than financial yield.
            </p>
          </section>

          {/* SECURITY & GOVERNANCE */}
          <section className="mb-6 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Security &amp; Governance
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Governance mechanisms and multi-sig frameworks protect critical
              parameters and institutional allocations. All changes are auditable
              and aligned with NOORCHAIN’s mission-driven, long-term vision.
            </p>
          </section>

          {/* ENGINEERING TRANSPARENCY */}
          <section className="p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm mb-6 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Engineering Transparency (Controlled)
            </h2>

            <p className="text-gray-700 leading-relaxed mb-4">
              NOORCHAIN is operated in a controlled, private mainnet-like environment.
              Public exposure is intentionally limited while feature completeness and
              security validation are in progress.
            </p>

            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <strong>Layer 1:</strong> sovereign EVM L1 with a permissioned BFT consensus layer.
              </li>
              <li>
                <strong>PoSS:</strong> application layer for governance and verifiable contribution signals (not consensus).
              </li>
              <li>
                <strong>Release artifacts:</strong> development is tracked through tagged milestones and versioned documentation.
              </li>
              <li>
                <strong>Public access posture:</strong> interfaces are opened progressively after review, not by default.
              </li>
            </ul>
          </section>

        </div>
      </div>
    </section>
  );
}