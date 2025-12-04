export default function TechnologyPage() {
  return (
    <main className="px-8 py-16 max-w-3xl">
      <h1 className="text-4xl font-bold mb-6">Technology</h1>

      <p className="text-lg text-gray-700 mb-6">
        NOORCHAIN is built using a modern blockchain architecture designed for
        security, clarity, and transparent social value creation.
      </p>

      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-2">Core Stack</h2>
        <p className="text-gray-700">
          The network is powered by a modular architecture including a social
          signal module (PoSS), a flexible account system, governance
          primitives, and an EVM-compatible execution layer.
        </p>
      </section>

      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-2">PoSS Integration</h2>
        <p className="text-gray-700">
          Proof of Signal Social is embedded directly into the chain logic,
          enabling transparent social participation and curator validation
          without financial incentives or yield promises.
        </p>
      </section>

      <section>
        <h2 className="text-2xl font-semibold mb-2">Security & Governance</h2>
        <p className="text-gray-700">
          Governance mechanisms ensure transparency and prevent unauthorized
          changes, while security rules protect the integrity of signals and
          validator operations.
        </p>
      </section>
    </main>
  );
}
