export default function LegalSection() {
  return (
    <section
      id="legal"
      className="relative w-full bg-transparent scroll-mt-[84px]"
    >
      <div className="relative z-10 container pb-10 md:pb-14">
        <div className="max-w-3xl">
          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white/10 border border-white/25 backdrop-blur-md px-3 py-1 mb-6">
            <span className="h-2 w-2 rounded-full bg-primary" />
            <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
              Legal & Compliance
            </span>
          </div>

          {/* TITLE */}
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            Legal & Compliance
          </h1>

          {/* INTRO */}
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/12 backdrop-blur-sm py-3 rounded-r-lg">
            This section summarizes the legal foundations, compliance rules, and
            risk disclosures governing the NOORCHAIN project. The framework follows
            a non-financial, transparency-first approach aligned with Legal Light CH.
          </p>

          {/* LEGAL FRAMEWORK */}
          <section
            id="legal-framework"
            className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 scroll-mt-24"
          >
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Legal Framework
            </h2>
            <p className="text-gray-700 leading-relaxed">
              NOORCHAIN operates under a compliance model that avoids investment
              solicitation, price promotion, or financial return promises.
              Protocol rules are transparent, auditable, and aligned with
              long-term public-interest use cases.
            </p>
          </section>

          {/* GOVERNANCE */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Governance &amp; Structure (Early Stage)
            </h2>

            <p className="text-gray-700 leading-relaxed mb-4">
              NOORCHAIN follows a separation-of-roles model designed to reduce
              conflicts of interest and support a Legal Light CH posture.
            </p>

            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li><strong>Protocol vs applications:</strong> infrastructure first.</li>
              <li><strong>Foundation (planned):</strong> Swiss non-profit stewardship.</li>
              <li><strong>Operational controls:</strong> auditable governance.</li>
              <li><strong>No financial services:</strong> no custody, no yield.</li>
            </ul>
          </section>

          {/* RISK DISCLOSURE */}
          <section
            id="risk-disclosure"
            className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 scroll-mt-24"
          >
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Risk Disclosure
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>Blockchain systems involve technical risks.</li>
              <li>Crypto-assets may be volatile.</li>
              <li>Regulatory frameworks may evolve.</li>
              <li>Users remain fully responsible for their actions.</li>
            </ul>
          </section>

          {/* NO INVESTMENT ADVICE */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              No Investment Advice
            </h2>
            <p className="text-gray-700 leading-relaxed">
              All content is informational only and does not constitute
              financial, legal, or investment advice.
            </p>
          </section>

          {/* LEGAL NOTICES */}
          <section
            id="legal-notices"
            className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5 scroll-mt-24"
          >
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Legal Notices
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Official terms, privacy rules, and intellectual property notices
              will be published here once finalized.
            </p>
          </section>

          {/* CONTACT */}
          <section className="p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm mb-10 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Contact
            </h2>
            <p className="text-gray-700 leading-relaxed">
              For governance or compliance matters:
              <br />
              <a
                href="mailto:contact@noorchain.io"
                className="text-primary hover:underline"
              >
                contact@noorchain.io
              </a>
            </p>
          </section>

        </div>
      </div>
    </section>
  );
}