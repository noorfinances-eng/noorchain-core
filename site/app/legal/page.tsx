export default function LegalPage() {
  return (
    <main className="px-8 py-16 max-w-3xl">
      <h1 className="text-4xl font-bold mb-6">Legal & Compliance</h1>

      <p className="text-lg text-gray-700 mb-10">
        This page summarizes the main legal principles, risk disclosures, and
        compliance considerations applicable to the NOORCHAIN project.
      </p>

      {/* Legal Framework */}
      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-3">Legal Framework</h2>
        <p className="text-gray-700">
          NOORCHAIN operates under a compliance model designed to avoid any form
          of financial solicitation, investment offering, or promise of returns.
          The protocol does not provide financial yield and does not promote
          token valuation.
        </p>
      </section>

      {/* Risk Disclosure */}
      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-3">Risk Disclosure</h2>
        <ul className="list-disc pl-6 text-gray-700 space-y-2">
          <li>Blockchain systems involve technical risks, including bugs or network instability.</li>
          <li>Cryptographic assets may be subject to strong volatility.</li>
          <li>Regulatory frameworks may evolve over time.</li>
          <li>Users remain fully responsible for their actions and interactions.</li>
        </ul>
      </section>

      {/* No investment advice */}
      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-3">No Investment Advice</h2>
        <p className="text-gray-700">
          All information provided is strictly educational and does not
          constitute investment advice, financial guidance, or any guarantee of
          performance.
        </p>
      </section>

      {/* Legal Notices */}
      <section className="mb-10">
        <h2 className="text-2xl font-semibold mb-3">Legal Notices</h2>
        <p className="text-gray-700">
          Terms of use, intellectual property, privacy rules, and operational
          disclaimers will be published here as official documents become
          available.
        </p>
      </section>

      {/* Contact */}
      <section>
        <h2 className="text-2xl font-semibold mb-3">Contact</h2>
        <p className="text-gray-700">
          For administrative, governance, or compliance-related matters, please
          contact us at:{" "}
          <a href="mailto:contact@noorchain.org" className="underline">
            contact@noorchain.org
          </a>
        </p>
      </section>
    </main>
  );
}
