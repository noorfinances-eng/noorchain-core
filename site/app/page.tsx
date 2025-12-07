export default function HomePage() {
  return (
    <main className="w-full">
      {/* HERO SECTION */}
      <section
        className="w-full"
        style={{
          backgroundImage: "url('/hero.svg')",
          backgroundSize: "cover",
          backgroundPosition: "center",
        }}
      >
        <div className="container py-24 md:py-32 text-white">
          <div className="max-w-3xl">
            <h1 className="text-4xl md:text-6xl font-bold leading-tight mb-6">
              NOORCHAIN
            </h1>

            <p className="text-lg md:text-xl text-white/90 mb-10">
              A Social Signal Blockchain powered by PoSS. Designed for transparent
              participation, curator validation, and a fixed-supply digital model
              free from financial speculation.
            </p>

            <div className="flex flex-wrap gap-4">
              <a
                href="/technology"
                className="px-6 py-3 bg-white text-gray-900 rounded-md text-sm font-medium hover:bg-gray-100"
              >
                Explore Technology
              </a>

              <a
                href="/genesis"
                className="px-6 py-3 border border-white text-white rounded-md text-sm font-medium hover:bg-white/10"
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
