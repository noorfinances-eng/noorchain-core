export default function PressPage() {
  return (
    <main className="w-full bg-paper">
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">

          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white border border-gray-soft px-3 py-1 mb-6">
            <span className="h-2 w-2 rounded-full bg-primary" />
            <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
              Press & Media
            </span>
          </div>

          {/* TITLE */}
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            NOORCHAIN Press & Media Kit
          </h1>

          {/* INTRO */}
          <p className="text-lg text-gray-700 leading-relaxed mb-8 border-l-4 border-primary pl-4 bg-white/60 py-3 rounded-r-lg">
            This page gathers verified information and official visual assets
            related to NOORCHAIN. Journalists, researchers, and institutional
            partners may use these materials for their work, provided that the
            brand guidelines and usage rules below are respected. NOORCHAIN
            does not provide investment advice and does not offer financial
            guarantees of any kind.
          </p>

          {/* ABOUT */}
          <section className="mb-10 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              About NOORCHAIN
            </h2>
            <p className="text-gray-700 leading-relaxed mb-4">
              NOORCHAIN is a Social Signal Blockchain designed to recognise and
              account for verified social participation. Rather than encouraging
              financial speculation, it focuses on meaningful actions validated
              by trusted curators such as NGOs, educators, and social actors.
              The protocol is built around a fixed supply, a transparent reward
              model, and a Legal Light framework that avoids investment
              solicitation and yield promises.
            </p>
            <h3 className="text-sm font-semibold uppercase tracking-wide text-gray-600 mb-2">
              Key Facts
            </h3>
            <ul className="list-disc pl-6 text-gray-700 space-y-1">
              <li>Type: Social Signal Blockchain (PoSS – Proof of Signal Social)</li>
              <li>Native asset: NUR (fixed supply)</li>
              <li>Supply cap: 299,792,458 NUR</li>
              <li>Reward split: 70% participant / 30% curator</li>
              <li>Halving: every 8 years</li>
              <li>Main focus: social participation, curator validation, transparent accounting</li>
              <li>Legal positioning: non-speculative, no financial promises, Legal Light approach</li>
            </ul>
          </section>

          {/* OFFICIAL DOCUMENTS */}
          <section className="mb-10 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Official Documents
            </h2>
            <p className="text-gray-700 leading-relaxed mb-4">
              These documents form the official reference for NOORCHAIN’s
              economic model, governance, genesis structure, and brand identity.
            </p>
            <ul className="list-disc pl-6 text-gray-700 space-y-2">
              <li>
                <span className="font-medium">NOORCHAIN Genesis Pack 1.1</span>{" "}
                — Genesis structure, allocation model 5/5/5/5/80, supply cap,
                halving, PoSS rules, governance foundations.{" "}
                <span className="text-xs text-gray-500">
                  (File: NOORCHAIN_Genesis_Pack_1.1.pdf)
                </span>
              </li>
              <li>
                <span className="font-medium">NOORCHAIN Brandbook 1.1</span>{" "}
                — Logo system, color palette, typography, banners, usage rules.{" "}
                <span className="text-xs text-gray-500">
                  (File: NOORCHAIN_Brandbook_1.1.pdf)
                </span>
              </li>
              <li>
                <span className="font-medium">
                  NOORCHAIN White Papers 1.1
                </span>{" "}
                — Investor, Storytelling, Public, and Long Version (20 pages),
                to be added once final layout is complete.
              </li>
            </ul>
          </section>

          {/* VISUAL ASSETS */}
          <section className="mb-10 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Visual Assets (Logos & Banners)
            </h2>
            <p className="text-gray-700 leading-relaxed mb-4">
              You may use the following visual assets when presenting or
              referencing NOORCHAIN. They must not be altered, recoloured
              outside the official palette, or used to promote financial
              products or unrelated services.
            </p>

            <h3 className="text-sm font-semibold uppercase tracking-wide text-gray-600 mb-2">
              Logos
            </h3>
            <ul className="list-disc pl-6 text-gray-700 space-y-1 mb-4">
              <li>Primary icon (N-frame symbol)</li>
              <li>Full logo (symbol + NOORCHAIN wordmark)</li>
              <li>Monochrome versions (navy, white)</li>
              <li>App and favicon icons</li>
            </ul>
            <p className="text-sm text-gray-600 mb-4">
              Recommended files (SVG):{" "}
              <span className="font-mono text-xs">
                logo-main.svg, logo-full.svg, logo-icon.svg, logo-mono-navy.svg,
                logo-mono-white.svg, appicon.svg, favicon.svg
              </span>
            </p>

            <h3 className="text-sm font-semibold uppercase tracking-wide text-gray-600 mb-2">
              Hero & Social Banners
            </h3>
            <ul className="list-disc pl-6 text-gray-700 space-y-1 mb-4">
              <li>Homepage hero banner</li>
              <li>X / Twitter banner</li>
              <li>LinkedIn banner</li>
              <li>GitHub banner</li>
            </ul>
            <p className="text-sm text-gray-600">
              Recommended files (SVG):{" "}
              <span className="font-mono text-xs">
                hero.svg, social-banner-x.svg, social-banner-linkedin.svg,
                social-banner-github.svg
              </span>
            </p>
          </section>

          {/* HOW TO PRESENT */}
          <section className="mb-10 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              How to Present NOORCHAIN
            </h2>

            <h3 className="text-sm font-semibold uppercase tracking-wide text-gray-600 mb-2">
              One-line description
            </h3>
            <p className="text-gray-700 leading-relaxed mb-4 italic">
              “NOORCHAIN is a Social Signal Blockchain that rewards verified
              social participation instead of financial speculation.”
            </p>

            <h3 className="text-sm font-semibold uppercase tracking-wide text-gray-600 mb-2">
              Short description
            </h3>
            <p className="text-gray-700 leading-relaxed mb-4">
              NOORCHAIN is a Social Signal Blockchain where value is linked to
              verified social actions rather than trading or speculation.
              Participants generate signals—such as micro-donations, verified
              participation, or certified content—and curators validate them
              according to transparent rules. Rewards follow a fixed, publicly
              documented model with a capped supply and without financial yield
              promises.
            </p>

            <p className="text-sm text-gray-600">
              Any description of NOORCHAIN should avoid suggesting guaranteed
              returns, price expectations, or investment advice. The project is
              positioned as a participation and utility framework, not as a
              speculative product.
            </p>
          </section>

          {/* LEGAL & USAGE RULES */}
          <section className="mb-10 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Legal & Usage Rules
            </h2>
            <p className="text-gray-700 leading-relaxed mb-4">
              NOORCHAIN does not encourage speculative trading and does not
              communicate price targets, performance expectations, or investment
              advice. All protocol rules are published to ensure transparency
              and governance clarity, and not as financial promotion.
            </p>
            <p className="text-gray-700 leading-relaxed mb-3">
              Partners and journalists are invited to respect this positioning
              and to avoid language implying guaranteed returns or investment
              products when referring to NOORCHAIN.
            </p>
            <h3 className="text-sm font-semibold uppercase tracking-wide text-gray-600 mb-2">
              When using NOORCHAIN brand assets:
            </h3>
            <ul className="list-disc pl-6 text-gray-700 space-y-1">
              <li>
                Do not imply that holding NUR guarantees a financial return.
              </li>
              <li>
                Do not use the logo to promote unrelated tokens, trading
                platforms, or financial instruments.
              </li>
              <li>Do not modify the official logo colours.</li>
              <li>Always refer clearly to the project as “NOORCHAIN”.</li>
            </ul>
          </section>

          {/* CONTACT */}
          <section className="mb-4 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Press & Institutional Contact
            </h2>
            <p className="text-gray-700 leading-relaxed mb-2">
              For press inquiries, partnerships, or institutional communication,
              please use the following contact channel:
            </p>
            <p className="text-gray-800 font-medium mb-2">
              Email:{" "}
              <a
                href="mailto:contact@noorchain.org"
                className="text-primary hover:underline"
              >
                contact@noorchain.org
              </a>
            </p>
            <p className="text-sm text-gray-600">
              A dedicated communication address for the future NOOR Foundation
              will be published once the foundation is formally established.
            </p>
          </section>

          {/* END LINE */}
          <div className="h-px w-full bg-gray-soft mt-6" />

        </div>
      </section>
    </main>
  );
}
