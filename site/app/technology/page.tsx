export default function TechnologyPage() {
  return (
    <main className="w-full">
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-gray-900 mb-6">
            Technology
          </h1>

          <p className="text-lg text-gray-700 leading-relaxed mb-10">
            NOORCHAIN is built using a clear, modular blockchain architecture
            designed for security, transparency, and verifiable social value.
          </p>

          <section className="mb-10">
            <h2 className="text-xl md:text-2xl font-semibold text-gray-900 mb-3">
              Core Stack
            </h2>
            <p className="text-gray-700 leading-relaxed">
              The network combines a robust core ledger, an EVM-compatible
              execution layer, and a dedicated social signal module (PoSS).
              This structure allows developers to build applications while the
              protocol enforces transparent rules for participation and
              accounting.
            </p>
          </section>

          <section className="mb-10">
            <h2 className="text-xl md:text-2xl font-semibold text-gray-900 mb-3">
            PoSS Integration
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Proof of Signal Social is integrated as a native module
              responsible for handling social signals, applying daily limits,
              and computing rewards according to public parameters. The logic
              focuses on participation and curatorship, not on financial yield.
            </p>
          </section>

          <section>
            <h2 className="text-xl md:text-2xl font-semibold text-gray-900 mb-3">
              Security &amp; Governance
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Governance mechanisms and multi-sig structures protect critical
              parameters and institutional allocations. Changes to the protocol
              are controlled, auditable, and aligned with the long-term mission
              of the project.
            </p>
          </section>
        </div>
      </section>
    </main>
  );
}
