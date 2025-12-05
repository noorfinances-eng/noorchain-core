export default function HomePage() {
  return (
    <main className="container py-20">
      <section className="max-w-3xl">
        <h1 className="text-6xl font-bold leading-tight mb-6">
          NOORCHAIN
        </h1>

        <p className="text-xl text-gray-700 mb-10">
          A Social Signal Blockchain powered by PoSS.  
          Designed for transparent participation, curator validation,  
          and a fixed-supply digital model free from financial speculation.
        </p>

        <div className="flex gap-4">
          <a
            href="/technology"
            className="px-6 py-3 bg-black text-white rounded-md text-sm hover:opacity-90"
          >
            Explore Technology
          </a>

          <a
            href="/genesis"
            className="px-6 py-3 border border-black rounded-md text-sm hover:bg-gray-100"
          >
            Genesis Overview
          </a>
        </div>
      </section>
    </main>
  );
}
