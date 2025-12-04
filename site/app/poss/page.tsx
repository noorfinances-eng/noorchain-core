export default function PoSSPage() {
  return (
    <main className="px-8 py-16 max-w-3xl">
      <h1 className="text-4xl font-bold mb-6">Proof of Signal Social (PoSS)</h1>

      <p className="text-lg text-gray-700 mb-6">
        PoSS is a participation-based mechanism designed to reward meaningful
        social signals without financial speculation or yield promises.
      </p>

      {/* Types of Signals */}
      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-3">Types of Social Signals</h2>
        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li>Micro-donations</li>
          <li>Verified participation</li>
          <li>Certified content</li>
          <li>CCN (Communication Network) signals</li>
        </ul>
      </section>

      {/* Reward Model */}
      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-3">Reward Distribution</h2>
        <p className="text-gray-700 mb-3">
          Rewards are allocated according to a clear fixed rule:
        </p>
        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li>70% to the participant generating the signal</li>
          <li>30% to the curator validating the signal</li>
        </ul>

        <p className="text-gray-700 mt-4">
          A halving mechanism applies every 8 years, and daily limits ensure
          fairness and protection against abuse.
        </p>
      </section>

      {/* Transparency */}
      <section>
        <h2 className="text-2xl font-semibold mb-3">Transparency & Parameters</h2>
        <p className="text-gray-700">
          All PoSS parameters are public and adjustable through governance.
          No financial guarantees or yield expectations are communicated or implied.
        </p>
      </section>
    </main>
  );
}
