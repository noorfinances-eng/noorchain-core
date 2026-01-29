export default function DocsSection() {
  return (
    <section
      id="docs"
      className="relative w-full bg-transparent scroll-mt-[84px]"
    >
      <div className="relative z-10 container pb-10 md:pb-14">
        <div className="max-w-3xl">
          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white/10 border border-white/25 backdrop-blur-md px-3 py-1 mb-6">
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
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/12 backdrop-blur-sm py-3 rounded-r-lg">
            Access the official public documents of NOORCHAIN, including white
            papers, technical specifications, governance rules, legal framework,
            and branding assets.
          </p>

          {/* BRANDING */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">Branding</h2>
            <p className="text-gray-700 leading-relaxed mb-3">
              The official visual identity of NOORCHAIN, including logo system,
              color palette, typography, and social assets.
            </p>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a
                  href="/docs/NOORCHAIN_Brandbook_1.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  NOORCHAIN Brandbook 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Press_Kit_1.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  NOORCHAIN Press Kit 1.1 (PDF)
                </a>
              </li>
            </ul>
          </section>

          {/* WHITE PAPERS */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              White Papers
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a
                  href="/docs/NOORCHAIN_Whitepaper_Investor_1.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  Investor Whitepaper 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Whitepaper_Storytelling_EN_1.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  Storytelling Whitepaper EN 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Whitepaper_Storytelling_FR_1.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  Storytelling Whitepaper FR 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Public_Whitepaper_1.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  Public Whitepaper 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/NOORCHAIN_Whitepaper_Long_1.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  Long Whitepaper 1.1 (PDF)
                </a>
              </li>
              <li>
                <a
                  href="/docs/Tokenomics_Public_1.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  Tokenomics Public 1.1 (PDF)
                </a>
              </li>
            </ul>
          </section>

          {/* PROTOCOL / RC NOTES */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Protocol — Release Candidate (RC)
            </h2>

            <p className="text-gray-700 leading-relaxed mb-3">
              Internal verification notes and reproducible evidence for the current
              RC baseline. These documents are not third-party security audits.
            </p>

            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a
                  href="/docs/M21_RC_VERIFICATION_NOTE_2.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  M21 RC Verification Note 2.1 (PDF)
                </a>
              </li>
            </ul>
          </section>

          {/* RESEARCH & POSITION PAPERS */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Research &amp; Position Papers
            </h2>

            <p className="text-gray-700 leading-relaxed mb-3">
              Selected research preprints and position papers. These documents are
              conceptual in scope and may be revised.
            </p>

            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <a
                  href="/docs/Beyond_State_Process_Time_Legitimacy_Preprint_v0.1.pdf"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  Beyond State: Process, Time, and Legitimacy in Blockchain Systems (Preprint v0.1 — PDF)
                </a>
                <div className="text-sm text-gray-700 mt-2">
                  <span className="italic">
                    Note: This preprint does not constitute a technical specification, roadmap, or investment proposal.
                  </span>
                </div>
              </li>
            </ul>
          </section>

          {/* DOWNLOAD */}
          <section className="p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm mb-10 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">Downloads</h2>
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
        </div>
      </div>
    </section>
  );
}
