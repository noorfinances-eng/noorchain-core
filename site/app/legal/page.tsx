export default function LegalPage() {
  return (
    <main className="container py-16">
      <h1 className="text-4xl font-bold mb-6">Legal & Compliance</h1>

      <p className="text-lg text-gray-700 mb-10 max-w-3xl">
        This section summarizes the main legal principles, compliance rules, and
        risk disclosures that apply to the NOORCHAIN project.
      </p>

      {/* Legal Framework */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Legal Framework</h2>
        <p className="text-gray-700">
          NOORCHAIN operates under a compliance model that avoids investment
          solicitation or financial return promises. All protocol rules are
          transparent and auditable.
        </p>
      </section>

      {/* Risk Disclosure */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Risk Disclosure</h2>
        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li>Blockchain systems involve operational and technical risks.</li>
          <li>Crypto-assets may experience significant volatility.</li>
          <li>Regulatory frameworks may evolve in the future.</li>
          <li>Users are fully responsible for their decisions and actions.</li>
        </ul>
      </section>

      {/* No Investment Advice */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">No Investment Advice</h2>
        <p className="text-gray-700">
          All content provided on this website is purely informational and does
          not constitute financial advice, legal advice, or a performance
          guarantee of any kind.
        </p>
      </section>

      {/* Legal Notices */}
      <section className="mb-10 max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Legal Notices</h2>
        <p className="text-gray-700">
          Official terms of use, privacy rules, and intellectual property
          notices will be published here as documents are finalized.
        </p>
      </section>

      {/* Contact */}
      <section className="max-w-3xl">
        <h2 className="text-2xl font-semibold mb-3">Contact</h2>
        <p className="text-gray-700">
          For governance, compliance, or administrative matters, please contact:{" "}
          <a href="mailto:contact@noorchain.org" className="underline">
            contact@noorchain.org
          </a>
        </p>
      </section>
    </main>
  );
}
