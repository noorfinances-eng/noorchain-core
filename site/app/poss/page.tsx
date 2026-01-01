export default function PoSSPage() {
  return (
    <main className="w-full bg-paper">
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">

          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white border border-gray-soft px-3 py-1 mb-6">
            <span className="h-2 w-2 rounded-full bg-primary" />
            <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
              Application Layer (Non-Consensus)
            </span>
          </div>

          {/* TITLE */}
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            Proof of Signal Social (PoSS)
          </h1>

          {/* INTRO */}
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/60 py-3 rounded-r-lg">
            PoSS is an application-layer mechanism for governance, coordination,
            and verifiable social signals. It is not the consensus layer: network
            security is provided by a permissioned BFT consensus. PoSS structures
            curator validation and participation signals without offering yield
            promises or custody.
          </p>

          {/* TYPES OF SIGNALS */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Types of Social Signals
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>Micro-donations</li>
              <li>Verified participation</li>
              <li>Certified content</li>
              <li>CCN (Communication Network) signals</li>
            </ul>
          </section>

          {/* REWARD DISTRIBUTION */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Reward Distribution
            </h2>

            <p className="text-gray-700 leading-relaxed mb-3">
              Rewards follow a strictly defined and transparent allocation rule:
            </p>

            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>70% to the participant generating the signal</li>
              <li>30% to the curator validating the signal</li>
            </ul>

            <p className="text-gray-700 leading-relaxed mt-4">
              PoSS includes an 8-year halving cycle and daily limits to ensure
              fairness, stability, and protection against abuse.
            </p>
          </section>

          {/* TRANSPARENCY */}
          <section className="p-6 border border-gray-soft rounded-xl bg-white shadow-sm mb-10">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Transparency & Parameters
            </h2>

            <p className="text-gray-700 leading-relaxed">
              All PoSS parameters are publicly visible and adjustable through
              governance. The mechanism provides no yield promises and remains
              aligned with the long-term mission of the project, ensuring
              accountability and auditability at every stage.
            </p>
          </section>

          {/* END LINE */}
          <div className="h-px w-full bg-gray-soft" />
        </div>
      </section>
    </main>
  );
}
