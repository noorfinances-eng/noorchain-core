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
        {/* On laisse juste une hauteur pour afficher le visuel */}
        <div className="py-24 md:py-32" />
      </section>

      {/* SECTION CONTENU — texte + boutons, comme avant */}
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">
          <h1 className="text-4xl md:text-6xl font-bold leading-tight mb-6">
            NOORCHAIN
          </h1>

          <p className="text-lg md:text-xl text-gray-700 mb-10">
            A Social Signal Blockchain powered by PoSS. Designed for transparent
            participation, curator validation, and a fixed-supply digital model
            free from financial speculation.
          </p>

          <div className="flex flex-wrap gap-4">
            <a
              href="/technology"
              className="px-6 py-3 bg-black text-white rounded-md text-sm md:text-base font-medium hover:opacity-90"
            >
              Explore Technology
            </a>

            <a
              href="/genesis"
              className="px-6 py-3 border border-black rounded-md text-sm md:text-base font-medium hover:bg-gray-100"
            >
              Genesis Overview
            </a>
          </div>
        </div>
      </section>
    </main>
  );
}
