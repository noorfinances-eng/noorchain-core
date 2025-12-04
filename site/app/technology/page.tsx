export default function TechnologyPage() {
  return (
    <main className="container py-16">
      <h1 className="text-4xl font-bold mb-6">Technology</h1>

      <p className="text-lg text-gray-700 mb-6">
        NOORCHAIN is built using a clear, modular blockchain architecture
        designed for security, transparency, and verifiable social value.
      </p>

      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-2">Core Stack</h2>
        <p className="text-gray-700">
          The network combines a robust core ledger, an EVM-compatible execution
          layer, and a dedicated social signal module (PoSS). This structure
          allows developers to build applications while the protocol enforces
          transparent rules for participation and accounting.
        </p>
      </section>

      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-2">PoSS Integration</h2>
        <p className="text-gray-700">
          Proof of Signal Social is integrated as a native module responsible
          for handling social signals, applying daily limits, and computing
          rewards according to public parameters. The logic focuses on
          participation and curatorship, not on financial yield.
        </p>
      </section>

      <section className="max-w-3xl">
        <h2 className="text-2xl font-semibold mb-2">Security & Governance</h2>
        <p className="text-gray-700">
          Governance mechanisms and multi-sig structures protect critical
          parameters and institutional allocations. Changes to the protocol are
          controlled, auditable, and aligned with the long-term mission of the
          project.
        </p>
      </section>
    </main>
  );
}
