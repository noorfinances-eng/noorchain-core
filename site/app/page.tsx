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

            {/* Testnet status — discreet */}
            <p className="text-xs sm:text-sm text-white/80 mb-4 font-medium">
              Testnet phase ongoing — controlled, experimental, non-financial.
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
                className="w-full sm:w-auto text-center px-6 py-3 bg-primary text-white rounded-md text-sm md:text-base font-medium hover:bg-blue-700 transition"
              >
                Explore Technology
              </a>

              <a
                href="/genesis"
                className="w-full sm:w-auto text-center px-6 py-3 border border-white text-white rounded-md text-sm md:text-base font-medium hover:bg-white/10 transition"
              >
                Genesis Overview
              </a>
            </div>
          </div>
        </div>
      </section>

      {/* SECTION ÉPURÉE — BLANCHE + TITRE + TEXTE */}
      <section className="container py-12 md:py-20">
        <div className="max-w-3xl">
          <h2 className="text-xl sm:text-2xl md:text-3xl font-bold text-gray-900 mb-4">
            A New Approach to Blockchain Design
          </h2>

          <p className="text-base sm:text-lg text-gray-700 leading-relaxed">
            NOORCHAIN introduces a mission-driven blockchain architecture focused
            on verified social participation rather than financial speculation.
            Powered by the PoSS protocol and aligned with Legal Light CH, it
            provides a transparent and sustainable digital infrastructure for
            curators, participants, institutions and communities.
          </p>
        </div>
      </section>
    </main>
  );
}
