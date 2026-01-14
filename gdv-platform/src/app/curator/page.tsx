"use client";

import { useMemo, useState } from "react";

function logout() {
  document.cookie = "gdv_role=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT; samesite=lax";
  localStorage.removeItem("gdv_role");
  window.location.href = "/login";
}


type Campaign = {
  id: string;
  name: string;
  period: string;
  rule: string;
  validatedActions: number;
  points: number;
  pendingCount: number;
};

type ActionRow = {
  id: string;
  anonId: string;
  actionType: string;
  qty: number;
  status: "À valider" | "Validé";
};

function cx(...parts: Array<string | false | null | undefined>) {
  return parts.filter(Boolean).join(" ");
}

function Icon({
  name,
  className,
}: {
  name:
    | "clock"
    | "user"
    | "menu"
    | "cloud"
    | "gear"
    | "plus"
    | "chevDown"
    | "pencil"
    | "check"
    | "sliders";
  className?: string;
}) {
  const c = className ?? "h-5 w-5";
  switch (name) {
    case "plus":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="M12 5v14M5 12h14" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" />
        </svg>
      );
    case "chevDown":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="m7 10 5 5 5-5" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" />
        </svg>
      );
    case "clock":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="M12 8v5l3 2" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" />
          <path d="M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" stroke="currentColor" strokeWidth="1.8" />
        </svg>
      );
    case "user":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="M20 21a8 8 0 1 0-16 0" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" />
          <path d="M12 13a4 4 0 1 0-4-4 4 4 0 0 0 4 4Z" stroke="currentColor" strokeWidth="1.8" />
        </svg>
      );
    case "menu":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="M5 7h14M5 12h14M5 17h14" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" />
        </svg>
      );
    case "cloud":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path
            d="M7 18a4 4 0 0 1 .6-7.95A5.5 5.5 0 0 1 18.6 12H19a3 3 0 0 1 0 6H7Z"
            stroke="currentColor"
            strokeWidth="1.8"
            strokeLinejoin="round"
          />
        </svg>
      );
    case "gear":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path
            d="M12 15.5a3.5 3.5 0 1 0-3.5-3.5 3.5 3.5 0 0 0 3.5 3.5Z"
            stroke="currentColor"
            strokeWidth="1.8"
          />
          <path
            d="M19.4 15a8.7 8.7 0 0 0 .1-2l2-1.2-2-3.4-2.3.7a8.5 8.5 0 0 0-1.7-1l-.3-2.4H10.8l-.3 2.4a8.5 8.5 0 0 0-1.7 1L6.5 8.4l-2 3.4 2 1.2a8.7 8.7 0 0 0 .1 2l-2 1.2 2 3.4 2.3-.7a8.5 8.5 0 0 0 1.7 1l.3 2.4h4.4l.3-2.4a8.5 8.5 0 0 0 1.7-1l2.3.7 2-3.4-2-1.2Z"
            stroke="currentColor"
            strokeWidth="1.2"
            strokeLinejoin="round"
            opacity="0.9"
          />
        </svg>
      );
    case "pencil":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path
            d="M4 20h4l10.5-10.5a2 2 0 0 0 0-3L16.5 4.5a2 2 0 0 0-3 0L3 15v5Z"
            stroke="currentColor"
            strokeWidth="1.6"
            strokeLinejoin="round"
          />
          <path d="M12.5 5.5 18.5 11.5" stroke="currentColor" strokeWidth="1.6" strokeLinecap="round" />
        </svg>
      );
    case "check":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="M20 7 10 17l-5-5" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" />
        </svg>
      );
    case "sliders":
      return (
        <svg className={c} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="M4 7h9M17 7h3" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" />
          <path d="M4 17h3M11 17h9" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" />
          <path d="M13 7a2 2 0 1 0-4 0 2 2 0 0 0 4 0Z" stroke="currentColor" strokeWidth="1.8" />
          <path d="M11 17a2 2 0 1 0-4 0 2 2 0 0 0 4 0Z" stroke="currentColor" strokeWidth="1.8" />
        </svg>
      );
  }
}

function GlassCard({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) {
  return (
    <div
      className={cx(
        "rounded-2xl border border-white/10 bg-white/[0.06] shadow-[0_8px_30px_rgba(0,0,0,0.35)] backdrop-blur-xl",
        className
      )}
    >
      {children}
    </div>
  );
}

export default function Home() {
  const campaigns = useMemo<Campaign[]>(() => [
  {
    id: "arbres-2024",
    name: "Plantons des arbres",
    period: "du 01 Mars au 01 Mai 2024",
    rule: "Planter au moins 10 arbres par semaine.",
    validatedActions: 150,
    points: 1500,
    pendingCount: 5,
  },
  {
    id: "education-ongoing",
    name: "Soutien éducatif",
    period: "campagne continue",
    rule: "Valider les actions éducatives hebdomadaires.",
    validatedActions: 42,
    points: 420,
    pendingCount: 2,
  },
], []);

  const rows = useMemo<ActionRow[]>(() => [
  { id: "row-1", anonId: "Élève-001", actionType: "Plantations d'arbres", qty: 12, status: "À valider" },
  { id: "row-2", anonId: "Élève-002", actionType: "Nettoyages de déchets", qty: 3, status: "À valider" },
  { id: "row-3", anonId: "Élève-003", actionType: "Reforestation", qty: 20, status: "À valider" },
  { id: "row-4", anonId: "Élève-001", actionType: "Plantations d'arbres", qty: 6, status: "À valider" },
  { id: "row-5", anonId: "Élève-006", actionType: "Plantations", qty: 5, status: "À valider" },
], []);

  const [activeCampaignId, setActiveCampaignId] = useState(campaigns[0]?.id ?? "");
  const [selected, setSelected] = useState<Record<string, boolean>>({
    "row-1": true,
    "row-3": true,
    "row-4": true,
  });

  const activeCampaign = useMemo(
    () => campaigns.find((c) => c.id === activeCampaignId) ?? campaigns[0],
    [activeCampaignId, campaigns]
  );

  const selectedCount = useMemo(() => Object.values(selected).filter(Boolean).length, [selected]);
  const allSelected = useMemo(() => rows.every((r) => selected[r.id]), [rows, selected]);

  function toggleRow(id: string) {
    setSelected((s) => ({ ...s, [id]: !s[id] }));
  }

  function toggleAll() {
    if (allSelected) {
      const next: Record<string, boolean> = {};
      rows.forEach((r) => (next[r.id] = false));
      setSelected(next);
      return;
    }
    const next: Record<string, boolean> = {};
    rows.forEach((r) => (next[r.id] = true));
    setSelected(next);
  }

  function clearSelection() {
    const next: Record<string, boolean> = {};
    rows.forEach((r) => (next[r.id] = false));
    setSelected(next);
  }

  return (
    <div className="min-h-screen bg-slate-950 text-slate-100">
      {/* background: dense institutional glass */}
      <div className="pointer-events-none fixed inset-0 -z-10">
        <div className="absolute inset-0 bg-[radial-gradient(1200px_700px_at_20%_10%,rgba(59,130,246,0.22),transparent_62%),radial-gradient(900px_540px_at_80%_18%,rgba(148,163,184,0.14),transparent_60%),radial-gradient(900px_520px_at_55%_85%,rgba(30,58,138,0.18),transparent_60%)]" />
        <div className="absolute inset-0 bg-gradient-to-b from-black/20 via-black/55 to-black/75" />
        <div className="absolute inset-0 opacity-[0.08] [background-image:radial-gradient(rgba(255,255,255,0.7)_1px,transparent_1px)] [background-size:22px_22px]" />
      </div>

      {/* Top bar */}
      <header className="sticky top-0 z-50 border-b border-white/10 bg-slate-950/55 backdrop-blur-xl">
        <div className="mx-auto flex max-w-7xl items-center justify-between px-6 py-4">
          <div className="flex items-center gap-3">
            <div className="flex h-9 w-9 items-center justify-center rounded-xl bg-white/10 ring-1 ring-white/15">
              <div className="h-4 w-4 rounded bg-blue-500/70" />
            </div>
            <div className="leading-tight">
              <div className="text-base font-semibold tracking-tight">
                Curators Hub <span className="text-slate-400">- Graine de Vie</span>
              </div>
            </div>
          </div>

          <div className="flex items-center gap-3">
            <button className="inline-flex items-center gap-2 rounded-xl bg-blue-500/20 px-4 py-2 text-sm font-medium text-blue-100 ring-1 ring-blue-400/25 hover:bg-blue-500/25">
              <Icon name="plus" className="h-4 w-4" />
              Nouvelle campagne
            </button>

            <div className="flex items-center gap-2 text-slate-300">
              <button className="rounded-xl bg-white/5 p-2 ring-1 ring-white/10 hover:bg-white/10" aria-label="Historique">
                <Icon name="clock" />
              </button>
              <button className="rounded-xl bg-white/5 p-2 ring-1 ring-white/10 hover:bg-white/10" aria-label="Compte">
                <Icon name="user" />
              </button>
              <button className="rounded-xl bg-white/5 p-2 ring-1 ring-white/10 hover:bg-white/10" aria-label="Menu">
                <Icon name="menu" />
              </button>
              <button className="rounded-xl bg-white/5 p-2 ring-1 ring-white/10 hover:bg-white/10" aria-label="Sync">
                <Icon name="cloud" />
              </button>
              <button className="rounded-xl bg-white/5 p-2 ring-1 ring-white/10 hover:bg-white/10" aria-label="Paramètres">
                <Icon name="gear" />
              </button>
            </div>
          </div>
        </div>
      </header>

      <main className="mx-auto max-w-7xl px-6 pb-28 pt-8">
      <div className="mx-auto max-w-6xl px-6 pt-8 pb-2 flex items-center justify-between">
        <div className="text-xs tracking-widest text-white/60">CURATOR · DASHBOARD</div>
        <div className="flex items-center gap-2">
          <a
            href="/curator"
            className="rounded-xl border border-white/25 bg-white/[0.10] px-4 py-2 text-sm text-white/90"
          >
            Dashboard
          </a>
          <a
            href="/curator/exports"
            className="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20"
          >
            Exports
          </a>
          <button
            type="button"
            onClick={logout}
            className="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20"
          >
            Se déconnecter
          </button>
        </div>
      </div>

        {/* Title */}
        <section>
          <h1 className="text-3xl font-semibold tracking-tight text-slate-100">
            Centre du curateur <span className="text-slate-400">- Graine de Vie</span>
          </h1>
          <p className="mt-2 max-w-3xl text-sm leading-6 text-slate-300/90">
            Validez les actions des élèves dans le cadre des campagnes en cours, puis envoyez-les à la blockchain pour les enregistrer.
          </p>
        </section>

        {/* Campaign selector band */}
        <section className="mt-6">
          <GlassCard className="px-4 py-4">
            <div className="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
              <div className="flex items-center gap-3">
                <span className="inline-flex h-7 w-7 items-center justify-center rounded-full bg-emerald-500/15 ring-1 ring-emerald-400/25">
                  <Icon name="check" className="h-4 w-4 text-emerald-200" />
                </span>
                <div className="text-sm text-slate-200">
                  <span className="font-medium">Campagne en cours :</span>{" "}
                  <span className="text-slate-300">{activeCampaign?.name}</span>{" "}
                  <span className="text-slate-400">- {activeCampaign?.period}</span>
                </div>
              </div>

              <div className="flex w-full items-center gap-2 md:w-[520px]">
                <div className="relative w-full">
                  <select
                    value={activeCampaignId}
                    onChange={(e) => setActiveCampaignId(e.target.value)}
                    className="h-11 w-full appearance-none rounded-xl border border-white/10 bg-slate-950/35 pl-3 pr-10 text-sm text-slate-100 ring-1 ring-white/5 focus:outline-none"
                  >
                    {campaigns.map((c) => (
                      <option key={c.id} value={c.id}>
                        {c.name} — {c.period}
                      </option>
                    ))}
                  </select>
                  <div className="pointer-events-none absolute right-3 top-1/2 -translate-y-1/2 text-slate-400">
                    <Icon name="chevDown" className="h-5 w-5" />
                  </div>
                </div>
              </div>
            </div>
          </GlassCard>
        </section>

        {/* Main grid */}
        <section className="mt-6 grid gap-6 lg:grid-cols-12">
          {/* Left: table */}
          <div className="lg:col-span-8">
            <GlassCard className="p-5">
              <div className="text-sm font-medium text-slate-200">
                Validez les actions anonymes effectuées par les élèves dans la campagne sélectionnée.
              </div>

              {/* toolbar */}
              <div className="mt-4 flex flex-wrap items-center gap-3">
                <button
                  onClick={toggleAll}
                  className="inline-flex items-center gap-2 rounded-xl bg-white/5 px-3 py-2 text-sm text-slate-200 ring-1 ring-white/10 hover:bg-white/10"
                >
                  <span className="inline-flex h-5 w-5 items-center justify-center rounded-md bg-white/10 ring-1 ring-white/15">
                    {allSelected ? "✓" : ""}
                  </span>
                  Tout sélectionner
                </button>

                <button
                  className="inline-flex items-center gap-2 rounded-xl bg-emerald-500/15 px-4 py-2 text-sm font-medium text-emerald-100 ring-1 ring-emerald-400/20 hover:bg-emerald-500/20"
                >
                  <Icon name="plus" className="h-4 w-4" />
                  Valider la sélection
                </button>

                <button
                  onClick={clearSelection}
                  className="inline-flex items-center gap-2 rounded-xl bg-white/5 px-4 py-2 text-sm text-slate-200 ring-1 ring-white/10 hover:bg-white/10"
                >
                  Vider la sélection
                </button>
              </div>

              {/* table */}
              <div className="mt-4 overflow-hidden rounded-2xl border border-white/10">
                <div className="flex items-center justify-between border-b border-white/10 bg-white/5 px-4 py-3">
                  <div className="text-xs font-medium text-slate-300"> </div>
                  <div className="flex items-center gap-2 text-slate-300">
                    <button className="rounded-lg bg-white/5 p-2 ring-1 ring-white/10 hover:bg-white/10" aria-label="Options">
                      <Icon name="sliders" className="h-4 w-4" />
                    </button>
                    <button className="rounded-lg bg-white/5 p-2 ring-1 ring-white/10 hover:bg-white/10" aria-label="Paramètres table">
                      <Icon name="gear" className="h-4 w-4" />
                    </button>
                  </div>
                </div>

                <table className="w-full text-left text-sm">
                  <thead className="bg-white/5 text-xs text-slate-300">
                    <tr>
                      <th className="w-12 px-4 py-3">
                        <button
                          onClick={toggleAll}
                          className="inline-flex h-5 w-5 items-center justify-center rounded-md bg-white/10 ring-1 ring-white/15"
                          aria-label="Tout sélectionner"
                        >
                          {allSelected ? "✓" : ""}
                        </button>
                      </th>
                      <th className="px-4 py-3 font-medium">ID Anonyme</th>
                      <th className="px-4 py-3 font-medium">Type d’action</th>
                      <th className="px-4 py-3 font-medium">Quantité</th>
                      <th className="px-4 py-3 font-medium">Statut</th>
                      <th className="px-4 py-3 text-right font-medium"> </th>
                    </tr>
                  </thead>

                  <tbody className="divide-y divide-white/10">
                    {rows.map((r) => {
                      const isSelected = !!selected[r.id];
                      return (
                        <tr key={r.id} className="bg-slate-950/10">
                          <td className="px-4 py-3">
                            <button
                              onClick={() => toggleRow(r.id)}
                              className={cx(
                                "inline-flex h-6 w-6 items-center justify-center rounded-md ring-1 transition",
                                isSelected
                                  ? "bg-emerald-500/15 text-emerald-200 ring-emerald-400/25"
                                  : "bg-white/5 text-transparent ring-white/15 hover:bg-white/10"
                              )}
                              aria-label={`Sélection ${r.anonId}`}
                            >
                              ✓
                            </button>
                          </td>
                          <td className="px-4 py-3 text-slate-200">{r.anonId}</td>
                          <td className="px-4 py-3 text-slate-300">{r.actionType}</td>
                          <td className="px-4 py-3 text-slate-200">{r.qty}</td>
                          <td className="px-4 py-3">
                            <span className="inline-flex items-center rounded-full bg-amber-500/15 px-3 py-1 text-xs font-medium text-amber-200 ring-1 ring-amber-500/20">
                              À valider
                            </span>
                          </td>
                          <td className="px-4 py-3">
                            <div className="flex justify-end gap-2">
                              <button className="rounded-lg bg-white/5 p-2 ring-1 ring-white/10 hover:bg-white/10" aria-label="Éditer">
                                <Icon name="pencil" className="h-4 w-4 text-slate-300" />
                              </button>
                              <button className="rounded-lg bg-white/5 p-2 ring-1 ring-white/10 hover:bg-white/10" aria-label="Valider">
                                <Icon name="check" className="h-4 w-4 text-slate-300" />
                              </button>
                            </div>
                          </td>
                        </tr>
                      );
                    })}
                  </tbody>
                </table>

                {/* pagination */}
                <div className="flex flex-wrap items-center justify-between gap-3 border-t border-white/10 bg-white/5 px-4 py-3 text-xs text-slate-300">
                  <div className="flex items-center gap-3">
                    <span>Page 1 sur 5</span>
                    <div className="flex items-center gap-2">
                      <span className="text-slate-400">10</span>
                      <Icon name="chevDown" className="h-4 w-4 text-slate-400" />
                    </div>
                  </div>

                  <div className="flex items-center gap-2">
                    <button className="rounded-lg bg-white/5 px-2 py-1 ring-1 ring-white/10 hover:bg-white/10" aria-label="Début">
                      |◀
                    </button>
                    <button className="rounded-lg bg-white/5 px-2 py-1 ring-1 ring-white/10 hover:bg-white/10" aria-label="Précédent">
                      ◀
                    </button>
                    <span className="text-slate-400">Page 1 sur 5</span>
                    <button className="rounded-lg bg-white/5 px-2 py-1 ring-1 ring-white/10 hover:bg-white/10" aria-label="Suivant">
                      ▶
                    </button>
                    <button className="rounded-lg bg-white/5 px-2 py-1 ring-1 ring-white/10 hover:bg-white/10" aria-label="Fin">
                      ▶|
                    </button>
                  </div>
                </div>
              </div>
            </GlassCard>
          </div>

          {/* Right: summary */}
          <div className="lg:col-span-4">
            <GlassCard className="p-5">
              <div className="text-lg font-semibold text-slate-200">Résumé de la campagne</div>

              <div className="mt-4 space-y-1 text-sm">
                <div className="text-slate-100">{activeCampaign?.name}</div>
                <div className="text-slate-400">{activeCampaign?.period}</div>
              </div>

              <div className="mt-4 text-sm text-slate-300">
                <span className="text-slate-400">Règle :</span> {activeCampaign?.rule}
              </div>

              <div className="mt-5 space-y-2 text-sm">
                <div className="text-slate-300">
                  <span className="text-slate-400">Total d’actions validées :</span> {activeCampaign?.validatedActions}
                </div>
                <div className="text-slate-300">
                  <span className="text-slate-400">Total points gagnés :</span> {activeCampaign?.points} points
                </div>
              </div>

              <button className="mt-6 w-full rounded-xl bg-blue-500/20 px-4 py-3 text-sm font-medium text-blue-100 ring-1 ring-blue-400/25 hover:bg-blue-500/25">
                Valider toutes les actions ({activeCampaign?.pendingCount ?? 0})
              </button>
            </GlassCard>
          </div>
        </section>

        {/* Bottom CTA like reference */}
        <div className="fixed bottom-6 right-6">
          <button
            className={cx(
              "rounded-xl px-6 py-3 text-sm font-semibold ring-1 transition",
              selectedCount > 0
                ? "bg-blue-500/20 text-blue-100 ring-blue-400/25 hover:bg-blue-500/25"
                : "bg-white/5 text-slate-400 ring-white/10"
            )}
            aria-disabled={selectedCount === 0}
          >
            Valider la sélection ({selectedCount})
          </button>
        </div>
      </main>
    </div>
  );
}
