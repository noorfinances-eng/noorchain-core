import Image from "next/image";

export default function HomePage() {
  return (
    <main className="w-full">
      {/* HERO SECTION — gradient + logo + tagline + texte + boutons */}
      <section className="w-full bg-gradient-to-br from-[#1A6AFF] to-[#00D1B2]">
        <div className="container py-16 md:py-20 text-white">
          <div className="max-w-3xl">
            {/* Logo + nom + tagline */}
            <div className="flex items-center gap-3 mb-4">
              <Image
                src="/logo-white.svg"
                alt="NOORCHAIN Logo"
                width={52}
                height={52}
                priority
              />
              <h1 className="text-3xl md:text-5xl font-extrabold tracking-tight">
                NOORCHAIN
              </h1>
            </div>

            {/* Tagline officielle */}
            <p className="text-lg md:text-xl text-white/90 mb-6 font-medium">
              A Human-Centered Blockchain for Social Signals
            </p>

            {/* Paragraphe explicatif */}
            <p className="text-base md:text-lg text-white/85 leading-relaxed mb-8">
              A Social Signal Blockchain powered by PoSS. Designed for transparent
              participation, curator validation, and a fixed-supply digital model
              free from financial speculation.
            </p>

            {/* Boutons — version premium NOORCHAIN */}
            <div className="flex flex-wrap gap-4">
              <a
                href="/technology"
                className="px-6 py-3 bg-primary text-white rounded-md text-sm md:text-base font-medium hover:bg-blue-700 transition"
              >
                Explore Technology
              </a>

              <a
                href="/genesis"
                className="px-6 py-3 border border-white text-white rounded-md text-sm md:text-base font-medium hover:bg-white/10 transition"
              >
                Genesis Overview
              </a>
            </div>
          </div>
        </div>
      </section>

      {/* SECTION ÉPURÉE — BLANCHE + TITRE + TEXTE (classe et sobre) */}
      <section className="container py-20">
        <div className="max-w-3xl">
          <h2 className="text-2xl md:text-3xl font-bold text-gray-900 mb-4">
            A New Approach to Blockchain Design
          </h2>

          <p className="text-gray-700 leading-relaxed text-lg">
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
