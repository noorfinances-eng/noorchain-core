export default function GenesisPage() {
  return (
    <main className="px-8 py-16 max-w-3xl">
      <h1 className="text-4xl font-bold mb-6">Genesis Overview</h1>

      <p className="text-lg text-gray-700 mb-6">
        The NOORCHAIN genesis defines the foundational rules of the network:
        fixed supply, halving schedule, PoSS distribution rules, and the
        allocation model that structures the chain’s long-term sustainability.
      </p>

      {/* Supply */}
      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-3">Total Supply</h2>
        <p className="text-gray-700">
          The total supply is fixed at <strong>299,792,458 NUR</strong>.  
          This supply can never increase.
        </p>
      </section>

      {/* Distribution */}
      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-3">Initial Distribution</h2>

        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li><strong>5%</strong> — NOOR Foundation</li>
          <li><strong>5%</strong> — Dev Pool</li>
          <li><strong>5%</strong> — PoSS Stimulus</li>
          <li><strong>5%</strong> — Optional Pre-sale</li>
          <li><strong>80%</strong> — PoSS Mintable Supply</li>
        </ul>
      </section>

      {/* Governance Principles */}
      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-3">Governance Principles</h2>
        <p className="text-gray-700">
          Genesis governance establishes the immutable foundations of the chain:
        </p>

        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li>Fixed supply cannot be changed.</li>
          <li>Halving occurs every 8 years.</li>
          <li>PoSS rewards always follow the 70/30 split.</li>
          <li>All parameters remain transparent and publicly auditable.</li>
        </ul>
      </section>

      {/* Downloads */}
      <section>
        <h2 className="text-2xl font-semibold mb-3">Genesis Pack</h2>
        <p className="text-gray-700 mb-3">
          The full genesis documents, including governance, allocation, and
          institutional structure, will be available here.
        </p>

        <a
          href="#"
          className="inline-block px-6 py-3 border border-black rounded-md text-sm"
        >
          Download Genesis Pack
        </a>
      </section>
    </main>
  );
}
