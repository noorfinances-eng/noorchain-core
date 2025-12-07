export default function TechnologyPage() {
  return (
    <main className="w-full">
      {/* Intro section légèrement mise en valeur */}
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            Technology
          </h1>

          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4">
            NOORCHAIN is built using a modular blockchain architecture 
            designed for security, transparency, and verifiable social value.
          </p>

          {/* Section 1 */}
          <section className="mb-12 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Core Stack
            </h2>
            <p className="text-gray-700 leading-relaxed">
              The network combines a core ledger, an EVM-compatible execution
              layer, and a dedicated PoSS module. This structure ensures 
              transparent participation rules while allowing developers to 
              build applications on top of a robust foundation.
            </p>
          </section>

          {/* Section 2 */}
          <section className="mb-12 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              PoSS Integration
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Proof of Signal Social is natively embedded into the protocol.
              It applies daily limits, computes rewards, and validates 
              participant–curator interactions according to public parameters.
            </p>
          </section>

          {/* Section 3 */}
          <section className="p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Security & Governance
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Governance mechanisms and multi-sig frameworks protect critical 
              parameters and institutional allocations. All changes are auditable 
              and aligned with NOORCHAIN’s mission-driven, long-term vision.
            </p>
          </section>
        </div>
      </section>
    </main>
  );
}
