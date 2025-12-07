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
              {/* Bouton principal : bleu NOORCHAIN */}
              <a
                href="/technology"
                className="px-6 py-3 bg-primary text-white rounded-md text-sm md:text-base font-medium hover:bg-blue-700 transition"
              >
                Explore Technology
              </a>

              {/* Bouton secondaire : bordure blanche */}
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

      {/* BANDEAU TURQUOISE — test de bg-light */}
      <section className="container py-8">
        <div className="bg-light text-white px-6 py-4 rounded-md">
          NOORCHAIN focuses on real-world social participation, curators, and
          mission-driven signals — not on speculative trading or financial promises.
        </div>
      </section>

      {/* SECTION PREMIUM — test de bg-indigo-deep */}
      <section className="container pb-16">
        <div className="bg-indigo-deep text-white py-12 px-6 rounded-xl">
          <h2 className="text-2xl md:text-3xl font-bold mb-4">
            What makes NOORCHAIN different?
          </h2>
          <p className="text-white/90 max-w-2xl leading-relaxed">
            NOORCHAIN is built as a Social Signal Blockchain: participants and
            curators are rewarded for verifiable social actions instead of
            energy-intensive mining. With a fixed supply, transparent governance
            and Legal Light compliance, NOORCHAIN is designed as a long-term public
            infrastructure rather than a short-term speculative asset.
          </p>
        </div>
      </section>
    </main>
  );
}
