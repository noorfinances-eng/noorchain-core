export default function HomePage() {
  return (
    <main className="w-full">
      {/* HERO SECTION — gradient + texte + boutons (sans logo à droite) */}
      <section className="w-full bg-gradient-to-br from-[#1A6AFF] to-[#00D1B2]">
        <div className="container py-16 md:py-20 text-white">
          <div className="max-w-3xl">
            <h1 className="text-3xl md:text-5xl font-extrabold tracking-tight mb-5">
              NOORCHAIN
            </h1>

            <p className="text-lg md:text-xl text-white/90 leading-relaxed mb-8">
              A Social Signal Blockchain powered by PoSS. Designed for transparent
              participation, curator validation, and a fixed-supply digital model
              free from financial speculation.
            </p>

            <div className="flex flex-wrap gap-4">
              {/* Bouton principal */}
              <a
                href="/technology"
                className="px-6 py-3 bg-white text-gray-900 rounded-md text-sm md:text-base font-medium hover:bg-gray-100 transition"
              >
                Explore Technology
              </a>

              {/* Bouton secondaire */}
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
    </main>
  );
}
