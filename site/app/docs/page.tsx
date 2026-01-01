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
            papers, technical specifications, governance rules, legal framework,
            and branding assets.
          </p>

          {/* BRANDING / BRANDBOOK */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Branding
            </h2>
            <p className="text-gray-700 leading-relaxed mb-3">
              The official visual identity of NOORCHAIN, including logo system,
              color palette, typography, and social assets.
            </p>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a
                  href="/docs/NOORCHAIN_Brandbook_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  NOORCHAIN Brandbook 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Press_Kit_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  NOORCHAIN Press Kit 1.1 (PDF)
                </a>
              </li>
            </ul>
          </section>

          {/* WHITE PAPERS */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              White Papers
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a
                  href="/docs/NOORCHAIN_Whitepaper_Investor_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Investor Whitepaper 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Whitepaper_Storytelling_EN_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Storytelling Whitepaper EN 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Whitepaper_Storytelling_FR_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Storytelling Whitepaper FR 1.1 (PDF)
                </a>
              </li>

              {/* >>> ADDED LINK HERE <<< */}
              <li>
                <a
                  href="/docs/NOORCHAIN_Public_Whitepaper_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Public Whitepaper 1.1 (PDF)
                </a>
              </li>
              {/* >>> END ADDITION <<< */}

              <li>
                <a
                  href="/docs/NOORCHAIN_Whitepaper_Long_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Long Whitepaper 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/Tokenomics_Public_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Tokenomics Public 1.1 (PDF)
                </a>
              </li>
            </ul>
          </section>

          {/* TECHNICAL DOCS */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Technical & Genesis Documents
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a
                  href="/docs/NOORCHAIN_Genesis_Pack_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  NOORCHAIN Genesis Pack 1.1 (PDF)
                </a>
              </li>

              {/* REMOVED: NOORCHAIN Testnet (Public 1.1, MD) */}

              <li>
                <a
                  href="/docs/NOORCHAIN_Genesis_Allocation_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Genesis Allocation 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Genesis_Parameters_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Genesis Parameters 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Genesis_Governance_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Genesis Governance 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Economic_Model_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Economic Model 1.1 (PDF)
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
                <a
                  href="/docs/Legal_Light_Framework_Public_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Legal Light Framework 2025 (Public 1.1, PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Legal_Architecture_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  NOORCHAIN Legal Architecture 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Foundation_Statutes_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Foundation Statutes 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Compliance_Framework_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Compliance Framework 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Multisig_Charter_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Multisig Committee Charter 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Legal_Notices_1.1.pdf"
                  target="_blank"
                  className="text-primary hover:underline"
                >
                  Legal Notices & Disclaimers 1.1 (PDF)
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
