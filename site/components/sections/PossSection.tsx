export default function PossSection() {
  return (
    <section
      id="poss"
      className="relative w-full bg-transparent scroll-mt-[84px]"
    >
      <div className="relative z-10 container pb-10 md:pb-14">
        <div className="max-w-3xl">
          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white/10 border border-white/25 backdrop-blur-md px-3 py-1 mb-6">
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
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/12 backdrop-blur-sm py-3 rounded-r-lg">
            PoSS is an application-layer mechanism for governance, coordination,
            and verifiable social signals. It is not the consensus layer: network
            security is provided by a permissioned BFT consensus.
          </p>

          {/* TYPES */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
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

          {/* REWARDS */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Reward Distribution
            </h2>

            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>70% to the participant generating the signal</li>
              <li>30% to the curator validating the signal</li>
            </ul>

            <p className="text-gray-700 leading-relaxed mt-4">
              PoSS includes an 8-year halving cycle and daily limits.
            </p>
          </section>

          {/* TRANSPARENCY */}
          <section className="p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Transparency & Parameters
            </h2>

            <p className="text-gray-700 leading-relaxed">
              All PoSS parameters are public, auditable, and governance-driven.
              No yield promises. No custody.
            </p>
          </section>

        </div>
      </div>
    </section>
  );
}