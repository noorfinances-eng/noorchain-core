export default function LegalPage() {
  return (
    <main className="w-full bg-paper">
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">

          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white border border-gray-soft px-3 py-1 mb-6">
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
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/60 py-3 rounded-r-lg">
            This section summarizes the legal foundations, compliance rules, and
            risk disclosures governing the NOORCHAIN project. The framework follows
            a non-financial, transparency-first approach aligned with Legal Light CH.
          </p>

          {/* LEGAL FRAMEWORK */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Legal Framework
            </h2>
            <p className="text-gray-700 leading-relaxed">
              NOORCHAIN operates under a compliance model that avoids investment
              solicitation, price promotion, or financial return promises. All
              protocol rules are transparent, auditable, and aligned with
              long-term public-interest use cases rather than speculation.
            </p>
          </section>

          {/* RISK DISCLOSURE */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Risk Disclosure
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>Blockchain systems involve operational and technical risks.</li>
              <li>Crypto-assets may experience significant volatility.</li>
              <li>Regulatory frameworks may evolve in the future.</li>
              <li>Users remain fully responsible for their decisions and actions.</li>
            </ul>
          </section>

          {/* NO INVESTMENT ADVICE */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              No Investment Advice
            </h2>
            <p className="text-gray-700 leading-relaxed">
              All content on this website is informational only. It does not
              constitute financial advice, legal advice, investment solicitation,
              or any form of performance guarantee.
            </p>
          </section>

          {/* LEGAL NOTICES */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Legal Notices
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Official terms of use, privacy rules, and intellectual property
              notices will be published here once finalized. All documents will
              remain accessible through the Documentation Hub.
            </p>
          </section>

          {/* CONTACT */}
          <section className="p-6 border border-gray-soft rounded-xl bg-white shadow-sm mb-10">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Contact
            </h2>
            <p className="text-gray-700 leading-relaxed">
              For governance, compliance, or administrative matters, please
              contact:{" "}
              <a
                href="mailto:contact@noorchain.io"
                className="text-primary hover:underline"
              >
                contact@noorchain.io
              </a>
            </p>
          </section>

          {/* END LINE */}
          <div className="h-px w-full bg-gray-soft" />
        </div>
      </section>
    </main>
  );
}
