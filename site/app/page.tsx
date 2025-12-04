export default function HomePage() {
  return (
    <main className="container py-16">
      {/* Hero Section */}
      <section className="max-w-3xl">
        <h1 className="text-5xl font-bold mb-4">NOORCHAIN</h1>

        <p className="text-xl text-gray-700 mb-6">
          A Social Signal Blockchain powered by PoSS, enabling transparent and
          ethical participation without financial speculation.
        </p>

        <div className="flex gap-4">
          <a
            href="/technology"
            className="px-6 py-3 bg-black text-white rounded-md text-sm"
          >
            Explore Technology
          </a>

          <a
            href="/genesis"
            className="px-6 py-3 border border-black rounded-md text-sm"
          >
            Genesis Overview
          </a>
        </div>
      </section>
    </main>
  );
}
