export default function DocsPage() {
  return (
    <main className="container py-16">
      <h1 className="text-4xl font-bold mb-6">Documentation</h1>

      <p className="text-lg text-gray-700 mb-10 max-w-3xl">
        Access the main public documents related to NOORCHAIN: white papers,
        technical specifications, governance rules, and legal framework.
      </p>

      {/* White Papers */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">White Papers</h2>
        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li><a href="#" className="underline">Investor White Paper</a></li>
          <li><a href="#" className="underline">Storytelling White Paper</a></li>
          <li><a href="#" className="underline">Public White Paper</a></li>
          <li><a href="#" className="underline">Long Version (20 pages)</a></li>
        </ul>
      </section>

      {/* Technical Docs */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Technical Documents</h2>
        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li><a href="#" className="underline">NOORCHAIN Master Summary 1.1</a></li>
          <li><a href="#" className="underline">Core Architecture & EVM Overview</a></li>
          <li><a href="#" className="underline">PoSS Module Documentation</a></li>
        </ul>
      </section>

      {/* Legal & Governance */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Legal & Governance</h2>
        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li><a href="#" className="underline">Legal Light Framework 2025</a></li>
          <li><a href="#" className="underline">Legal Architecture</a></li>
          <li><a href="#" className="underline">Foundation Statutes</a></li>
          <li><a href="#" className="underline">Governance Charter</a></li>
          <li><a href="#" className="underline">MultiSig Committee Rules</a></li>
        </ul>
      </section>

      {/* Downloads */}
      <section className="max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Downloads</h2>
        <p className="text-gray-700 mb-3">
          Final documents and archives will be available in downloadable format
          (PDF or ZIP) from this section.
        </p>
        <a
          href="#"
          className="inline-block px-6 py-3 border border-black rounded-md text-sm"
        >
          Download All Docs (coming soon)
        </a>
      </section>
    </main>
  );
}
