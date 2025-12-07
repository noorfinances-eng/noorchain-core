export default function DocsPage() {
  return (
    <main className="w-full bg-paper">
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">

          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white border border-gray-soft px-3 py-1 mb-6">
            <span className="h-2 w-2 rounded-full bg-primary" />
            <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
              Documentation Hub
            </span>
          </div>

          {/* TITLE */}
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            Documentation
          </h1>

          {/* INTRO */}
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/60 py-3 rounded-r-lg">
            Access the official public documents of NOORCHAIN, including white
            papers, technical specifications, governance rules, and legal
            framework.
          </p>

          {/* WHITE PAPERS */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              White Papers
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a href="#" className="text-primary hover:underline">
                  Investor White Paper
                </a>
              </li>
              <li>
                <a href="#" className="text-primary hover:underline">
                  Storytelling White Paper
                </a>
              </li>
              <li>
                <a href="#" className="text-primary hover:underline">
                  Public White Paper
                </a>
              </li>
              <li>
                <a href="#" className="text-primary hover:underline">
                  Long Version (20 pages)
                </a>
              </li>
            </ul>
          </section>

          {/* TECHNICAL DOCS */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Technical Documents
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a href="#" className="text-primary hover:underline">
                  NOORCHAIN Master Summary 1.1
                </a>
              </li>
              <li>
                <a href="#" className="text-primary hover:underline">
                  Core Architecture & EVM Overview
                </a>
              </li>
              <li>
                <a href="#" className="text-primary hover:underline">
                  PoSS Module Documentation
                </a>
              </li>
            </ul>
          </section>

          {/* LEGAL & GOVERNANCE */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Legal & Governance
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a href="#" className="text-primary hover:underline">
                  Legal Light Framework 2025
                </a>
              </li>
              <li>
                <a href="#" className="text-primary hover:underline">
                  Legal Architecture
                </a>
              </li>
              <li>
                <a href="#" className="text-primary hover:underline">
                  Foundation Statutes
                </a>
              </li>
              <li>
                <a href="#" className="text-primary hover:underline">
                  Governance Charter
                </a>
              </li>
              <li>
                <a href="#" className="text-primary hover:underline">
                  MultiSig Committee Rules
                </a>
              </li>
            </ul>
          </section>

          {/* DOWNLOAD */}
          <section className="p-6 border border-gray-soft rounded-xl bg-white shadow-sm mb-10">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Downloads
            </h2>
            <p className="text-gray-700 leading-relaxed mb-4">
              Final documents and archives will be available in downloadable
              format (PDF or ZIP) from this section.
            </p>
            <a
              href="#"
              className="inline-block px-6 py-3 border border-primary text-primary rounded-md text-sm md:text-base font-medium hover:bg-primary hover:text-white transition"
            >
              Download All Docs (coming soon)
            </a>
          </section>

          {/* END LINE */}
          <div className="h-px w-full bg-gray-soft" />
        </div>
      </section>
    </main>
  );
}
