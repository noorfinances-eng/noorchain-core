export default function CuratorsPage() {
  return (
    <main className="container py-16">
      <h1 className="text-4xl font-bold mb-6">Curators</h1>

      <p className="text-lg text-gray-700 mb-10 max-w-3xl">
        Curators play a key role in maintaining the integrity of PoSS by
        validating legitimate social signals. They are mission-driven partners,
        not financial actors.
      </p>

      {/* Who are the Curators */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Who are the Curators?</h2>
        <p className="text-gray-700">
          Curators include community organizations, NGOs, educators, and trusted
          social actors who contribute to fair validation and transparent signal
          assessment.
        </p>
      </section>

      {/* Curator Levels */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Curator Levels</h2>
        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li><strong>Bronze</strong> — Basic visibility and validation rights.</li>
          <li><strong>Silver</strong> — Enhanced visibility and participation scope.</li>
          <li><strong>Gold</strong> — Highest recognition and extended validation responsibility.</li>
        </ul>
      </section>

      {/* Recognition */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Recognition & Contribution</h2>
        <p className="text-gray-700">
          Curators earn recognition through their contribution to social
          integrity. They may receive PoSS-generated rewards (30% share) but
          without financial guarantees or yield expectations.
        </p>
      </section>

      {/* Becoming a Curator */}
      <section className="max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Becoming a Curator</h2>
        <p className="text-gray-700">
          A future application process will allow qualified organizations to
          join the curator network. Requirements focus on mission alignment,
          transparency, and social impact.
        </p>
      </section>
    </main>
  );
}
