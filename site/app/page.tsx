export default function HomePage() {
  return (
    <main className="w-full">

      {/* HERO SECTION — image seulement */}
      <section
        className="w-full"
        style={{
          backgroundImage: "url('/hero.svg')",
          backgroundSize: "cover",
          backgroundPosition: "center",
        }}
      >
        {/* Hauteur réduite pour un rendu plus sobre */}
        <div className="py-12 md:py-20" />
      </section>

      {/* SECTION CONTENU — texte + boutons */}
      <section className="container py-14 md:py-16">
        <div className="max-w-3xl">

          <h1 className="text-3xl md:text-5xl font-extrabold tracking-tight mb-5 text-gray-900">
            NOORCHAIN
          </h1>

          <p className="text-lg md:text-xl text-gray-600 leading-relaxed mb-10">
            A Social Signal Blockchain powered by PoSS. Designed for transparent
            participation, curator validation, and a fixed-supply digital model
            free from financial speculation.
          </p>

          <div className="flex flex-wrap gap-4">

            {/* Bouton principal — premium blue */}
            <a
              href="/technology"
              className="px-6 py-3 bg-[#1A6AFF] text-white rounded-md text-sm md:text-base font-medium hover:bg-blue-700 transition"
            >
              Explore Technology
            </a>

            {/* Bouton secondaire — border blue */}
            <a
              href="/genesis"
              className="px-6 py-3 border border-[#1A6AFF] text-[#1A6AFF] rounded-md text-sm md:text-base font-medium hover:bg-blue-50 transition"
            >
              Genesis Overview
            </a>

          </div>
        </div>
      </section>

    </main>
  );
}
