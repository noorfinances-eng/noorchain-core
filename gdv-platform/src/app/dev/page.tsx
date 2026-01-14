"use client";

import { useMemo, useState } from "react";

type ExportStatus = "Queued" | "Generating" | "Ready" | "Downloaded" | "Failed";

type ExportRow = {
  id: string;
  createdAtIso: string;
  campaign: string;
  preset: "Evidence Pack" | "Audit Summary" | "Raw Actions";
  format: "JSON" | "PDF";
  status: ExportStatus;
};

type ActionRow = {
  id: string;
  tsIso: string;
  actor: string;
  action: string;
  result: "OK" | "WARN" | "FAIL";
};

function cx(...parts: Array<string | false | null | undefined>) {
  return parts.filter(Boolean).join(" ");
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

function Pill({ children }: { children: React.ReactNode }) {
  return (
    <span className="inline-flex items-center rounded-full border border-white/10 bg-white/[0.06] px-3 py-1 text-xs text-white/80">
      {children}
    </span>
  );
}

function logout() {
  document.cookie = "gdv_role=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT; samesite=lax";
  localStorage.removeItem("gdv_role");
  window.location.href = "/login";
}

function formatDate(iso: string) {
  const d = new Date(iso);
  return d.toLocaleString(undefined, { year: "numeric", month: "short", day: "2-digit", hour: "2-digit", minute: "2-digit" });
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

function resultBadge(r: ActionRow["result"]) {
  const base = "inline-flex items-center rounded-full border px-2.5 py-1 text-xs";
  if (r === "OK") return <span className={cx(base, "border-emerald-400/25 bg-emerald-400/10 text-emerald-100")}>OK</span>;
  if (r === "WARN") return <span className={cx(base, "border-amber-400/25 bg-amber-400/10 text-amber-100")}>WARN</span>;
  return <span className={cx(base, "border-red-400/25 bg-red-400/10 text-red-100")}>FAIL</span>;
}

export default function DevPage() {
  const exportsRows = useMemo<ExportRow[]>(
    () => [
      {
        id: "exp-003",
        createdAtIso: "2026-01-14T17:10:00.000Z",
        campaign: "Plantons des arbres",
        preset: "Evidence Pack",
        format: "JSON",
        status: "Ready",
      },
      {
        id: "exp-002",
        createdAtIso: "2026-01-14T16:33:00.000Z",
        campaign: "Soutien éducatif",
        preset: "Audit Summary",
        format: "JSON",
        status: "Downloaded",
      },
      {
        id: "exp-001",
        createdAtIso: "2026-01-14T16:00:00.000Z",
        campaign: "Plantons des arbres",
        preset: "Raw Actions",
        format: "JSON",
        status: "Downloaded",
      },
    ],
    []
  );

  const actions = useMemo<ActionRow[]>(
    () => [
      { id: "a-01", tsIso: "2026-01-14T17:12:00.000Z", actor: "dev", action: "RBAC middleware enabled", result: "OK" },
      { id: "a-02", tsIso: "2026-01-14T16:58:00.000Z", actor: "curator", action: "Generate export exp-003", result: "OK" },
      { id: "a-03", tsIso: "2026-01-14T16:22:00.000Z", actor: "dev", action: "Update UI exports builder", result: "OK" },
    ],
    []
  );

  const [note, setNote] = useState<string>("");

  return (
    <main className="min-h-screen bg-[#071025] text-white">
      <div className="mx-auto max-w-6xl px-6 py-10 space-y-6">
        {/* Header */}
        <div className="flex items-start justify-between gap-4">
          <div>
            <div className="text-xs tracking-widest text-white/60">DEV · CONSOLE</div>
            <h1 className="mt-2 text-2xl font-semibold">Dev Console</h1>
            <div className="mt-3 flex flex-wrap items-center gap-2">
              <Pill>RBAC ON</Pill>
              <Pill>Prototype</Pill>
              <Pill>Next.js 16.1.1</Pill>
            </div>
          </div>

          <div className="hidden sm:flex items-center gap-2">
            <a
              href="/curator"
              className="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20"
            >
              Curator Dashboard
            </a>
            <a
              href="/curator/exports"
              className="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20"
            >
              Curator Exports
            </a>
            <a
              href="/dev"
              className="rounded-xl border border-white/25 bg-white/[0.10] px-4 py-2 text-sm text-white/90"
            >
              Dev
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

        {/* Top metrics */}
        <div className="grid grid-cols-1 gap-6 md:grid-cols-12">
          <Card className="md:col-span-4">
            <div className="p-6">
              <div className="text-xs tracking-widest text-white/60">SYSTEM</div>
              <div className="mt-2 text-lg font-semibold">Local Dev</div>
              <div className="mt-2 text-sm text-white/70">Server: http://localhost:3001 (webpack)</div>
              <div className="mt-4 grid grid-cols-2 gap-2 text-xs text-white/65">
                <div className="rounded-xl border border-white/10 bg-white/[0.04] px-3 py-2">Lint: OK</div>
                <div className="rounded-xl border border-white/10 bg-white/[0.04] px-3 py-2">RBAC: ON</div>
                <div className="rounded-xl border border-white/10 bg-white/[0.04] px-3 py-2">Routes: OK</div>
                <div className="rounded-xl border border-white/10 bg-white/[0.04] px-3 py-2">Session: cookie</div>
              </div>
            </div>
          </Card>

          <Card className="md:col-span-4">
            <div className="p-6">
              <div className="text-xs tracking-widest text-white/60">EXPORTS</div>
              <div className="mt-2 text-lg font-semibold">{exportsRows.length}</div>
              <div className="mt-2 text-sm text-white/70">Exports récents (mock data)</div>
              <div className="mt-4 text-xs text-white/65">
                Dernier: <span className="font-mono text-white/85">{exportsRows[0]?.id}</span> · {statusBadge(exportsRows[0]?.status ?? "Queued")}
              </div>
            </div>
          </Card>

          <Card className="md:col-span-4">
            <div className="p-6">
              <div className="text-xs tracking-widest text-white/60">CONTROLS</div>
              <div className="mt-2 text-lg font-semibold">Quick Actions</div>
              <div className="mt-4 space-y-2">
                <button
                  type="button"
                  onClick={() => setNote("Seed: done (prototype).")}
                  className="w-full rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20 text-left"
                >
                  Seed demo data
                </button>
                <button
                  type="button"
                  onClick={() => setNote("Session cleared (prototype).")}
                  className="w-full rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20 text-left"
                >
                  Clear local session
                </button>
                <button
                  type="button"
                  onClick={() => setNote("Export pipeline simulated (prototype).")}
                  className="w-full rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20 text-left"
                >
                  Simulate export pipeline
                </button>
              </div>

              {note ? (
                <div className="mt-3 rounded-xl border border-white/10 bg-white/[0.04] px-3 py-2 text-xs text-white/70">
                  {note}
                </div>
              ) : null}
            </div>
          </Card>
        </div>

        {/* Tables */}
        <div className="grid grid-cols-1 gap-6 lg:grid-cols-12">
          <Card className="lg:col-span-7">
            <div className="p-6">
              <div className="text-xs tracking-widest text-white/60">RECENT EXPORTS</div>
              <h2 className="mt-2 text-lg font-semibold">Exports</h2>

              <div className="mt-5 overflow-hidden rounded-2xl border border-white/10">
                <div className="grid grid-cols-12 bg-white/[0.04] px-4 py-3 text-xs text-white/60">
                  <div className="col-span-3">ID</div>
                  <div className="col-span-3">Campaign</div>
                  <div className="col-span-2">Preset</div>
                  <div className="col-span-2">Created</div>
                  <div className="col-span-2 text-right">Status</div>
                </div>

                <div className="divide-y divide-white/10">
                  {exportsRows.map((r) => (
                    <div key={r.id} className="grid grid-cols-12 items-center px-4 py-3 text-sm">
                      <div className="col-span-3 font-mono text-xs text-white/85">{r.id}</div>
                      <div className="col-span-3 text-white/85">{r.campaign}</div>
                      <div className="col-span-2 text-white/85">{r.preset}</div>
                      <div className="col-span-2 text-xs text-white/65">{formatDate(r.createdAtIso)}</div>
                      <div className="col-span-2 flex justify-end">{statusBadge(r.status)}</div>
                    </div>
                  ))}
                </div>
              </div>

              <div className="mt-4 text-xs text-white/50">
                Cette console est une vue Dev (prototype). Les données réelles seront branchées ensuite (exports persistés, preuves, logs).
              </div>
            </div>
          </Card>

          <Card className="lg:col-span-5">
            <div className="p-6">
              <div className="text-xs tracking-widest text-white/60">AUDIT TRAIL</div>
              <h2 className="mt-2 text-lg font-semibold">Actions</h2>

              <div className="mt-5 overflow-hidden rounded-2xl border border-white/10">
                <div className="grid grid-cols-12 bg-white/[0.04] px-4 py-3 text-xs text-white/60">
                  <div className="col-span-4">Time</div>
                  <div className="col-span-2">Actor</div>
                  <div className="col-span-4">Action</div>
                  <div className="col-span-2 text-right">Result</div>
                </div>

                <div className="divide-y divide-white/10">
                  {actions.map((a) => (
                    <div key={a.id} className="grid grid-cols-12 items-center px-4 py-3 text-sm">
                      <div className="col-span-4 text-xs text-white/65">{formatDate(a.tsIso)}</div>
                      <div className="col-span-2 text-white/85">{a.actor}</div>
                      <div className="col-span-4 text-white/85">{a.action}</div>
                      <div className="col-span-2 flex justify-end">{resultBadge(a.result)}</div>
                    </div>
                  ))}
                </div>
              </div>

              <div className="mt-4 flex flex-wrap gap-2">
                <a
                  href="/curator"
                  className="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20"
                >
                  Go to Curator
                </a>
                <a
                  href="/curator/exports"
                  className="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20"
                >
                  Go to Exports
                </a>
                <button
                  type="button"
                  onClick={logout}
                  className="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20"
                >
                  Logout
                </button>
              </div>
            </div>
          </Card>
        </div>
      </div>
    </main>
  );
}
