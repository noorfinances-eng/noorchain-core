"use client";

import { useMemo, useState } from "react";

import { CuratorTopBar } from "../../../components/CuratorTopBar";
function logout() {
  document.cookie = "gdv_role=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT; samesite=lax";
  localStorage.removeItem("gdv_role");
  window.location.href = "/login";
}

type ExportStatus = "Queued" | "Generating" | "Ready" | "Downloaded" | "Failed";

type Campaign = {
  id: string;
  name: string;
  period: string;
  scope: string;
};

type ExportPreset = "Evidence Pack" | "Audit Summary" | "Raw Actions";

type ExportRow = {
  id: string;
  createdAtIso: string;
  campaignId: string;
  preset: ExportPreset;
  range: "Last 7 days" | "Last 30 days" | "Custom";
  from?: string;
  to?: string;
  include: {
    actions: boolean;
    signatures: boolean;
    snapshotMeta: boolean;
    receipts: boolean;
    attachments: boolean;
  };
  format: "JSON" | "PDF";
  status: ExportStatus;
};

function cx(...parts: Array<string | false | null | undefined>) {
  return parts.filter(Boolean).join(" ");
}

function Pill({ children }: { children: React.ReactNode }) {
  return (
    <span className="inline-flex items-center rounded-full border border-white/10 bg-white/[0.06] px-3 py-1 text-xs text-white/80">
      {children}
    </span>
  );
}

function Card({ children, className }: { children: React.ReactNode; className?: string }) {
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

function formatDate(iso: string) {
  const d = new Date(iso);
  return d.toLocaleString(undefined, { year: "numeric", month: "short", day: "2-digit", hour: "2-digit", minute: "2-digit" });
}

function downloadJson(filename: string, payload: unknown) {
  const blob = new Blob([JSON.stringify(payload, null, 2)], { type: "application/json;charset=utf-8" });
  const url = URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = filename;
  document.body.appendChild(a);
  a.click();
  a.remove();
  URL.revokeObjectURL(url);
}

export default function CuratorExportsPage() {
  const campaigns = useMemo<Campaign[]>(
    () => [
      { id: "arbres-2024", name: "Plantons des arbres", period: "01 Mar 2024 → 01 May 2024", scope: "Graine de Vie (GDV)" },
      { id: "education-ongoing", name: "Soutien éducatif", period: "Campagne continue", scope: "GDV · Éducation" },
    ],
    []
  );

  const presets = useMemo<Array<{ id: ExportPreset; label: string; desc: string }>>(
    () => [
      { id: "Evidence Pack", label: "Evidence Pack", desc: "Pack complet pour preuve, archivage et traçabilité." },
      { id: "Audit Summary", label: "Audit Summary", desc: "Résumé exécutable (KPIs, compte-rendu, anomalies)." },
      { id: "Raw Actions", label: "Raw Actions", desc: "Export brut des actions (ingestion/analytics)." },
    ],
    []
  );

  const [campaignId, setCampaignId] = useState<string>(campaigns[0]?.id ?? "");
  const [preset, setPreset] = useState<ExportPreset>("Evidence Pack");

  const [range, setRange] = useState<ExportRow["range"]>("Last 7 days");
  const [from, setFrom] = useState<string>("");
  const [to, setTo] = useState<string>("");

  const [format, setFormat] = useState<ExportRow["format"]>("JSON");

  const [includeActions, setIncludeActions] = useState(true);
  const [includeSignatures, setIncludeSignatures] = useState(true);
  const [includeSnapshotMeta, setIncludeSnapshotMeta] = useState(true);
  const [includeReceipts, setIncludeReceipts] = useState(true);
  const [includeAttachments, setIncludeAttachments] = useState(false);

  const [rows, setRows] = useState<ExportRow[]>(() => [
    {
      id: "exp-001",
      createdAtIso: new Date(Date.now() - 1000 * 60 * 55).toISOString(),
      campaignId: "arbres-2024",
      preset: "Evidence Pack",
      range: "Last 7 days",
      include: { actions: true, signatures: true, snapshotMeta: true, receipts: true, attachments: false },
      format: "JSON",
      status: "Ready",
    },
    {
      id: "exp-002",
      createdAtIso: new Date(Date.now() - 1000 * 60 * 15).toISOString(),
      campaignId: "education-ongoing",
      preset: "Audit Summary",
      range: "Last 30 days",
      include: { actions: true, signatures: true, snapshotMeta: true, receipts: false, attachments: false },
      format: "JSON",
      status: "Downloaded",
    },
  ]);

  const activeCampaign = useMemo(() => campaigns.find((c) => c.id === campaignId) ?? campaigns[0], [campaignId, campaigns]);

  const canGenerate = useMemo(() => {
    if (!campaignId) return false;
    if (range === "Custom") {
      if (!from || !to) return false;
      const a = new Date(from).getTime();
      const b = new Date(to).getTime();
      if (!Number.isFinite(a) || !Number.isFinite(b)) return false;
      if (a > b) return false;
    }
    return true;
  }, [campaignId, range, from, to]);

  function nextId() {
    const n = rows.length + 1;
    return `exp-${String(n).padStart(3, "0")}`;
  }

  function generate() {
    if (!canGenerate) return;

    const id = nextId();
    const createdAtIso = new Date().toISOString();

    const newRow: ExportRow = {
      id,
      createdAtIso,
      campaignId,
      preset,
      range,
      from: range === "Custom" ? from : undefined,
      to: range === "Custom" ? to : undefined,
      include: {
        actions: includeActions,
        signatures: includeSignatures,
        snapshotMeta: includeSnapshotMeta,
        receipts: includeReceipts,
        attachments: includeAttachments,
      },
      format,
      status: "Queued",
    };

    setRows((r) => [newRow, ...r]);

    // Prototype: simulate generation pipeline
    window.setTimeout(() => {
      setRows((r) =>
        r.map((x) => (x.id === id ? { ...x, status: "Generating" } : x))
      );
    }, 450);

    window.setTimeout(() => {
      setRows((r) =>
        r.map((x) => (x.id === id ? { ...x, status: "Ready" } : x))
      );
    }, 1100);
  }

  function exportRow(row: ExportRow) {
    // Prototype JSON payload; ready for later binding to on-chain proofs / receipts.
    const payload = {
      meta: {
        product: "GDV-PLATFORM",
        role: "curator",
        exportedAtIso: new Date().toISOString(),
        exportId: row.id,
        format: row.format,
        preset: row.preset,
      },
      campaign: campaigns.find((c) => c.id === row.campaignId) ?? null,
      range: row.range === "Custom" ? { kind: row.range, from: row.from, to: row.to } : { kind: row.range },
      include: row.include,
      proof: {
        // placeholders for future on-chain proofs
        chainId: "TBD",
        snapshotId: "TBD",
        txHash: "TBD",
        receipts: row.include.receipts ? "TBD" : "omitted",
        signatures: row.include.signatures ? "TBD" : "omitted",
      },
      data: {
        actions: row.include.actions ? [] : "omitted",
        attachments: row.include.attachments ? [] : "omitted",
      },
    };

    const fname = `${row.id}-${row.campaignId}-${row.preset.replace(/\s+/g, "_")}.json`;
    downloadJson(fname, payload);

    setRows((r) => r.map((x) => (x.id === row.id ? { ...x, status: "Downloaded" } : x)));
  }

  function statusBadge(s: ExportStatus) {
    const base = "inline-flex items-center rounded-full border px-2.5 py-1 text-xs";
    switch (s) {
      case "Queued":
        return <span className={cx(base, "border-white/10 bg-white/[0.06] text-white/75")}>Queued</span>;
      case "Generating":
        return <span className={cx(base, "border-white/10 bg-white/[0.06] text-white/90")}>Generating</span>;
      case "Ready":
        return <span className={cx(base, "border-emerald-400/25 bg-emerald-400/10 text-emerald-100")}>Ready</span>;
      case "Downloaded":
        return <span className={cx(base, "border-sky-400/25 bg-sky-400/10 text-sky-100")}>Downloaded</span>;
      case "Failed":
        return <span className={cx(base, "border-red-400/25 bg-red-400/10 text-red-100")}>Failed</span>;
    }
  }
return (
    <main className="min-h-screen bg-[#071025] text-white">
      <div className="mx-auto max-w-6xl px-6 py-10 space-y-6">

        {/* Header */}
          <CuratorTopBar active="exports" onLogout={logout} />

          <section className="pt-6">
            <div className="text-xs tracking-widest text-white/60">CURATOR · EXPORTS</div>
            <h1 className="mt-2 text-2xl font-semibold">Export Evidence Pack</h1>
            <div className="mt-3 flex flex-wrap items-center gap-2">
              <Pill>{activeCampaign?.scope ?? "GDV"}</Pill>
              <Pill>Proof-ready (prototype)</Pill>
              <Pill>Institutional format</Pill>
            </div>
          </section>


        <div className="grid grid-cols-1 gap-6 lg:grid-cols-12">

          {/* Left: Builder */}
          <Card className="lg:col-span-5">
            <div className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <div className="text-xs tracking-widest text-white/60">EXPORT BUILDER</div>
                  <h2 className="mt-2 text-lg font-semibold">Configuration</h2>
                  <p className="mt-2 text-sm text-white/70">
                    Prépare un pack exportable pour preuve, archivage, ou audit.
                  </p>
                </div>
              </div>

              <div className="mt-6 space-y-5">
                {/* Campaign */}
                <div>
                  <label className="text-sm text-white/80">Campagne</label>
                  <select
                    value={campaignId}
                    onChange={(e) => setCampaignId(e.target.value)}
                    className="mt-2 w-full rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm outline-none focus:border-white/25"
                  >
                    {campaigns.map((c) => (
                      <option key={c.id} value={c.id} className="bg-[#071025]">
                        {c.name} — {c.period}
                      </option>
                    ))}
                  </select>
                </div>

                {/* Preset */}
                <div>
                  <label className="text-sm text-white/80">Preset</label>
                  <div className="mt-2 grid grid-cols-1 gap-2 sm:grid-cols-3">
                    {presets.map((p) => {
                      const active = preset === p.id;
                      return (
                        <button
                          key={p.id}
                          type="button"
                          onClick={() => setPreset(p.id)}
                          className={cx(
                            "rounded-xl border px-4 py-3 text-left",
                            active ? "border-white/25 bg-white/[0.10]" : "border-white/10 bg-white/[0.06] hover:border-white/20"
                          )}
                        >
                          <div className="text-sm font-semibold">{p.label}</div>
                          <div className="mt-1 text-xs text-white/65">{p.desc}</div>
                        </button>
                      );
                    })}
                  </div>
                </div>

                {/* Range */}
                <div>
                  <label className="text-sm text-white/80">Période</label>
                  <div className="mt-2 grid grid-cols-1 gap-2 sm:grid-cols-3">
                    {(["Last 7 days", "Last 30 days", "Custom"] as const).map((r) => {
                      const active = range === r;
                      return (
                        <button
                          key={r}
                          type="button"
                          onClick={() => setRange(r)}
                          className={cx(
                            "rounded-xl border px-4 py-3 text-left text-sm",
                            active ? "border-white/25 bg-white/[0.10]" : "border-white/10 bg-white/[0.06] hover:border-white/20"
                          )}
                        >
                          {r}
                        </button>
                      );
                    })}
                  </div>

                  {range === "Custom" ? (
                    <div className="mt-3 grid grid-cols-1 gap-2 sm:grid-cols-2">
                      <div>
                        <label className="text-xs text-white/65">From</label>
                        <input
                          value={from}
                          onChange={(e) => setFrom(e.target.value)}
                          type="date"
                          className="mt-2 w-full rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm outline-none focus:border-white/25"
                        />
                      </div>
                      <div>
                        <label className="text-xs text-white/65">To</label>
                        <input
                          value={to}
                          onChange={(e) => setTo(e.target.value)}
                          type="date"
                          className="mt-2 w-full rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm outline-none focus:border-white/25"
                        />
                      </div>
                    </div>
                  ) : null}
                </div>

                {/* Include */}
                <div>
                  <label className="text-sm text-white/80">Inclure</label>
                  <div className="mt-2 grid grid-cols-1 gap-2 sm:grid-cols-2">
                    <label className="flex items-center justify-between rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm hover:border-white/20">
                      <span className="text-white/85">Actions</span>
                      <input type="checkbox" checked={includeActions} onChange={(e) => setIncludeActions(e.target.checked)} />
                    </label>

                    <label className="flex items-center justify-between rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm hover:border-white/20">
                      <span className="text-white/85">Signatures curator</span>
                      <input
                        type="checkbox"
                        checked={includeSignatures}
                        onChange={(e) => setIncludeSignatures(e.target.checked)}
                      />
                    </label>

                    <label className="flex items-center justify-between rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm hover:border-white/20">
                      <span className="text-white/85">Snapshot metadata</span>
                      <input
                        type="checkbox"
                        checked={includeSnapshotMeta}
                        onChange={(e) => setIncludeSnapshotMeta(e.target.checked)}
                      />
                    </label>

                    <label className="flex items-center justify-between rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm hover:border-white/20">
                      <span className="text-white/85">Receipts (tx)</span>
                      <input type="checkbox" checked={includeReceipts} onChange={(e) => setIncludeReceipts(e.target.checked)} />
                    </label>

                    <label className="flex items-center justify-between rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm hover:border-white/20 sm:col-span-2">
                      <span className="text-white/85">Attachments (photos, PDFs)</span>
                      <input
                        type="checkbox"
                        checked={includeAttachments}
                        onChange={(e) => setIncludeAttachments(e.target.checked)}
                      />
                    </label>
                  </div>

                  <div className="mt-2 text-xs text-white/55">
                    En prototype, “Receipts / Signatures / Snapshot” sont des placeholders prêts à être connectés à NOORCHAIN.
                  </div>
                </div>

                {/* Format + CTA */}
                <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
                  <div className="flex items-center gap-2">
                    <label className="text-sm text-white/80">Format</label>
                    <select
                      value={format}
                      onChange={(e) => setFormat(e.target.value as ExportRow["format"])}
                      className="rounded-xl border border-white/10 bg-white/[0.06] px-3 py-2 text-sm outline-none focus:border-white/25"
                    >
                      <option value="JSON" className="bg-[#071025]">
                        JSON
                      </option>
                      <option value="PDF" className="bg-[#071025]" disabled>
                        PDF (next)
                      </option>
                    </select>
                  </div>

                  <button
                    type="button"
                    onClick={generate}
                    disabled={!canGenerate}
                    className="rounded-xl bg-white px-4 py-3 text-sm font-semibold text-[#071025] disabled:opacity-60"
                  >
                    Generate Export
                  </button>
                </div>
              </div>
            </div>
          </Card>

          {/* Right: History */}
          <Card className="lg:col-span-7">
            <div className="p-6">
              <div className="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
                <div>
                  <div className="text-xs tracking-widest text-white/60">EXPORT HISTORY</div>
                  <h2 className="mt-2 text-lg font-semibold">Exports récents</h2>
                  <p className="mt-2 text-sm text-white/70">
                    Générations et téléchargements. Chaque export devient un artefact archivable.
                  </p>
                </div>

                <div className="text-xs text-white/55">
                  {rows.length} export{rows.length > 1 ? "s" : ""}
                </div>
              </div>

              <div className="mt-6 overflow-hidden rounded-2xl border border-white/10">
                <div className="grid grid-cols-12 bg-white/[0.04] px-4 py-3 text-xs text-white/60">
                  <div className="col-span-3">ID</div>
                  <div className="col-span-3">Campaign</div>
                  <div className="col-span-2">Preset</div>
                  <div className="col-span-2">Created</div>
                  <div className="col-span-2 text-right">Status / Action</div>
                </div>

                <div className="divide-y divide-white/10">
                  {rows.map((r) => {
                    const camp = campaigns.find((c) => c.id === r.campaignId);
                    const canDownload = r.status === "Ready" || r.status === "Downloaded";
                    return (
                      <div key={r.id} className="grid grid-cols-12 items-center px-4 py-3 text-sm">
                        <div className="col-span-3 font-mono text-xs text-white/85">{r.id}</div>
                        <div className="col-span-3">
                          <div className="text-sm">{camp?.name ?? r.campaignId}</div>
                          <div className="text-xs text-white/55">{r.range === "Custom" ? "Custom range" : r.range}</div>
                        </div>
                        <div className="col-span-2 text-white/85">{r.preset}</div>
                        <div className="col-span-2 text-xs text-white/65">{formatDate(r.createdAtIso)}</div>
                        <div className="col-span-2 flex items-center justify-end gap-2">
                          {statusBadge(r.status)}
                          <button
                            type="button"
                            onClick={() => exportRow(r)}
                            disabled={!canDownload}
                            className={cx(
                              "rounded-xl border px-3 py-2 text-xs",
                              canDownload
                                ? "border-white/10 bg-white/[0.06] text-white/85 hover:border-white/20"
                                : "border-white/10 bg-white/[0.03] text-white/40"
                            )}
                          >
                            Download
                          </button>
                        </div>
                      </div>
                    );
                  })}
                </div>
              </div>

              <div className="mt-4 text-xs text-white/50">
                Note: “Download” génère un fichier JSON local. La version PDF sera ajoutée après stabilisation (et/ou génération côté serveur).
              </div>
            </div>
          </Card>
        </div>
      </div>
    </main>
  );
}
