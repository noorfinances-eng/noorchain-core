export default function TechnologyPage() {
  return (
    <main className="w-full bg-paper">
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">

          {/* Petit label de contexte */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white border border-gray-soft px-3 py-1 mb-6">
            <span className="h-2 w-2 rounded-full bg-primary" />
            <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
              Core Architecture
            </span>
          </div>

          {/* Titre principal */}
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            Technology
          </h1>

          {/* Paragraphe d’intro avec accent gauche */}
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/60 py-3 rounded-r-lg">
            NOORCHAIN is built using a modular blockchain architecture designed
            for security, transparency, and verifiable social value.
          </p>

          {/* Section 1 */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Core Stack
            </h2>
            <p className="text-gray-700 leading-relaxed">
              The network combines a core ledger, an EVM-compatible execution
              layer, and a dedicated PoSS module. This structure ensures
              transparent participation rules while allowing developers to build
              applications on top of a robust foundation.
            </p>
          </section>

          {/* Section 2 */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              PoSS Integration
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Proof of Signal Social is natively embedded into the protocol. It
              applies daily limits, computes rewards, and validates
              participant–curator interactions according to public parameters,
              focusing on participation and curatorship rather than financial
              yield.
            </p>
          </section>

          {/* Section 3 */}
          <section className="p-6 border border-gray-soft rounded-xl bg-white shadow-sm mb-10">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Security &amp; Governance
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Governance mechanisms and multi-sig frameworks protect critical
              parameters and institutional allocations. All changes are auditable
              and aligned with NOORCHAIN’s mission-driven, long-term vision.
            </p>
          </section>

          {/* Ligne de fin subtile pour donner l’impression de “page finie” */}
          <div className="h-px w-full bg-gray-soft" />
        </div>
      </section>
    </main>
  );
}
