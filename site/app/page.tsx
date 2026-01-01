import Image from "next/image";

export default function HomePage() {
  return (
    <main className="w-full">
      {/* HERO SECTION — gradient + logo + tagline + texte + boutons */}
      <section className="w-full bg-gradient-to-br from-[#1A6AFF] to-[#00D1B2]">
        <div className="container py-12 sm:py-16 md:py-20 text-white">
          <div className="max-w-3xl">
            {/* Logo + nom + tagline */}
            <div className="flex flex-col sm:flex-row items-start sm:items-center gap-2 sm:gap-3 mb-4">
              <Image
                src="/logo-white.svg"
                alt="NOORCHAIN Logo"
                width={52}
                height={52}
                priority
              />
              <h1 className="text-2xl sm:text-3xl md:text-5xl font-extrabold tracking-tight">
                NOORCHAIN
              </h1>
            </div>

            {/* Tagline officielle */}
            <p className="text-base sm:text-lg md:text-xl text-white/90 mb-2 font-medium">
              A Human-Centered Blockchain for Social Signals
            </p>

            {/* Project status — discreet */}
            <p className="text-xs sm:text-sm text-white/80 mb-4 font-medium">
              Private mainnet-like environment — controlled operation, non-public
              by design.
            </p>

            {/* Paragraphe explicatif */}
            <p className="text-sm sm:text-base md:text-lg text-white/85 leading-relaxed mb-6 sm:mb-8">
              A Social Signal Blockchain powered by PoSS. Designed for transparent
              participation, curator validation, and a fixed-supply digital model
              free from financial speculation.
            </p>

            {/* Boutons — version premium NOORCHAIN */}
            <div className="flex flex-col sm:flex-row flex-wrap gap-3 sm:gap-4">
              <a
                href="/technology"
                className="w-full sm:w-auto text-center px-6 py-3 bg-primary text-white rounded-md text-sm md:text-base font-medium hover:bg-blue-700 shadow-sm hover:shadow-md hover:-translate-y-0.5 active:translate-y-0 transition-all duration-200 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white/80"
              >
                Explore Technology
              </a>

              <a
                href="/genesis"
                className="w-full sm:w-auto text-center px-6 py-3 border border-white text-white rounded-md text-sm md:text-base font-medium hover:bg-white/10 shadow-sm hover:shadow-md hover:-translate-y-0.5 active:translate-y-0 transition-all duration-200 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white/80"
              >
                Genesis Overview
              </a>
            </div>
          </div>
        </div>
      </section>

      {/* STATUS + INTRO + PoSS FRAMING — grouped for consistent spacing and visual rhythm */}
      <section className="container py-10 md:py-14">
        <div className="max-w-3xl space-y-6">
          {/* CURRENT PROJECT STATUS — factual, dated */}
          <div className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-4">
              <h2 className="text-lg sm:text-xl font-bold text-gray-900">
                Current Project Status
              </h2>
              <p className="text-xs sm:text-sm text-gray-500">
                Last updated: 2026-01-01
              </p>
            </div>

            <ul className="text-sm sm:text-base text-gray-700 leading-relaxed space-y-2">
              <li>
                <span className="font-semibold text-gray-900">Network:</span>{" "}
                Private mainnet-like environment (continuous operation)
              </li>
              <li>
                <span className="font-semibold text-gray-900">Consensus:</span>{" "}
                Permissioned BFT
              </li>
              <li>
                <span className="font-semibold text-gray-900">Layer 1:</span>{" "}
                Sovereign EVM L1
              </li>
              <li>
                <span className="font-semibold text-gray-900">PoSS:</span>{" "}
                Application layer for governance and verifiable contribution
                signals (not consensus)
              </li>
              <li>
                <span className="font-semibold text-gray-900">Public access:</span>{" "}
                Limited by design until feature completeness and security review
              </li>
              <li>
                <span className="font-semibold text-gray-900">Reference build:</span>{" "}
                M10-MAINNETLIKE-STABLE / M11-DAPPS-STABLE / M12.2-WORLDSTATE-RPC-NONCE-BALANCE
              </li>
            </ul>
          </div>

          {/* INTRO — boxed, same dimensions as status */}
          <div className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-lg sm:text-xl font-bold text-gray-900 mb-4">
              A New Approach to Blockchain Design
            </h2>

            <p className="text-sm sm:text-base text-gray-700 leading-relaxed">
              NOORCHAIN introduces a mission-driven blockchain architecture
              focused on verified social participation rather than financial
              speculation. Powered by the PoSS protocol and aligned with Legal
              Light CH, it provides a transparent and sustainable digital
              infrastructure for curators, participants, institutions and
              communities.
            </p>
          </div>

          {/* PoSS framing — minimal, home-level, links to /poss */}
          <div className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-4">
              <h2 className="text-lg sm:text-xl font-bold text-gray-900">
                PoSS Framing (Non-Consensus)
              </h2>
              <a
                href="/poss"
                className="text-sm font-medium text-blue-700 hover:text-blue-900 transition"
              >
                Read PoSS framing
              </a>
            </div>

            <ul className="text-sm sm:text-base text-gray-700 leading-relaxed space-y-2">
              <li>
                <span className="font-semibold text-gray-900">PoSS is not consensus.</span>{" "}
                Network security is provided by a permissioned BFT consensus layer.
              </li>
              <li>
                <span className="font-semibold text-gray-900">PoSS is an application layer.</span>{" "}
                It structures governance, coordination, and verifiable contribution signals.
              </li>
              <li>
                <span className="font-semibold text-gray-900">Economic posture:</span>{" "}
                NUR is the native gas token; NOORCHAIN does not offer returns or custody.
              </li>
            </ul>
          </div>
        </div>
      </section>
    </main>
  );
}
