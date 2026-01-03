export default function CuratorsSection() {
  return (
    <section
      id="curators"
      className="relative w-full bg-transparent scroll-mt-[84px]"
    >
      <div className="container pb-10 md:pb-14 relative z-10">
        <div className="max-w-3xl">
          {/* LABEL */}
          <div className="inline-flex items-center gap-2 rounded-full bg-white/10 border border-white/25 backdrop-blur-md px-3 py-1 mb-6">
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
          <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/12 backdrop-blur-sm py-3 rounded-r-lg">
            Curators are mission-driven actors responsible for validating
            legitimate social signals within NOORCHAIN. They protect the
            integrity of PoSS and ensure that participation remains meaningful,
            fair, and transparent.
          </p>

          {/* WHO ARE THE CURATORS */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
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
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
            <h2 className="text-2xl font-semibold text-navy mb-3">
              Curator Levels
            </h2>
            <ul className="list-disc pl-6 text-gray-700 leading-relaxed space-y-2">
              <li>
                <strong>Bronze</strong> — Basic visibility and limited validation rights.
              </li>
              <li>
                <strong>Silver</strong> — Enhanced visibility and increased validation responsibilities.
              </li>
              <li>
                <strong>Gold</strong> — Highest recognition with extended validation scope and institutional trust.
              </li>
            </ul>
          </section>

          {/* RECOGNITION */}
          <section className="mb-8 p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
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
          <section className="p-6 border border-white/20 rounded-xl bg-white/10 backdrop-blur-md shadow-sm mb-10 transition-all duration-200 hover:shadow-md hover:-translate-y-0.5">
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

        </div>
      </div>
    </section>
  );
}