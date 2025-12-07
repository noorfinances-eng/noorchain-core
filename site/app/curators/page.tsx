export default function CuratorsPage() {
  return (
    <main className="w-full bg-paper">
      <section className="container py-16 md:py-20">
        <div className="max-w-3xl">

          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white border border-gray-soft px-3 py-1 mb-6">
            <span className="h-2 w-2 rounded-full bg-primary" />
            <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
              Social Validation Layer
            </span>
          </div>

          {/* TITLE */}
          <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
            Curators
          </h1>

          {/* INTRO */}
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/60 py-3 rounded-r-lg">
            Curators are mission-driven actors responsible for validating
            legitimate social signals within NOORCHAIN. They protect the
            integrity of PoSS and ensure that participation remains meaningful,
            fair, and transparent.
          </p>

          {/* WHO ARE THE CURATORS */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Who are the Curators?
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Curators include community organizations, NGOs, educators, and
              trusted social institutions. They act as verifiers of real
              participation, helping ensure that PoSS signals reflect genuine
              social contribution rather than automated or manipulative behavior.
            </p>
          </section>

          {/* LEVELS */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Curator Levels
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <strong>Bronze</strong> — Basic visibility and limited validation
                rights.
              </li>
              <li>
                <strong>Silver</strong> — Enhanced visibility and increased
                validation responsibilities.
              </li>
              <li>
                <strong>Gold</strong> — Highest recognition with extended
                validation scope and institutional trust.
              </li>
            </ul>
          </section>

          {/* RECOGNITION */}
          <section className="mb-8 p-6 border border-gray-soft rounded-xl bg-white shadow-sm">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Recognition & Contribution
            </h2>
            <p className="text-gray-700 leading-relaxed">
              Curators receive social recognition and increased visibility for
              their contribution to ecosystem integrity. They may earn PoSS
              rewards (the 30% curator share), but without any financial
              guarantees or yield expectations — ensuring compliance with Legal
              Light CH.
            </p>
          </section>

          {/* BECOMING A CURATOR */}
          <section className="p-6 border border-gray-soft rounded-xl bg-white shadow-sm mb-10">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Becoming a Curator
            </h2>
            <p className="text-gray-700 leading-relaxed">
              A structured application process will allow qualified organizations
              to join the curator network. Requirements will emphasize mission
              alignment, transparency, trusted identity, and measurable social
              impact.
            </p>
          </section>

          {/* END LINE */}
          <div className="h-px w-full bg-gray-soft" />
        </div>
      </section>
    </main>
  );
}
