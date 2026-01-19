"use client";

import React from "react";

type Active = "dashboard" | "exports";

function cx(...parts: Array<string | false | null | undefined>) {
  return parts.filter(Boolean).join(" ");
}

function PlusIcon({ className }: { className?: string }) {
  const base = cx("h-5 w-5", className);
  return (
    <svg className={base} viewBox="0 0 24 24" fill="none" aria-hidden="true">
      <path d="M12 5v14M5 12h14" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
    </svg>
  );
}

export function CuratorTopBar({
  active,
  onLogout,
  onNewCampaign,
  subtitle = "Graine de Vie",
  orgLabel = "Curators Hub",
}: {
  active: Active;
  onLogout: () => void;
  onNewCampaign?: () => void;
  subtitle?: string;
  orgLabel?: string;
}) {
  return (
    <header className="sticky top-0 z-50 border-b border-white/10 bg-slate-950/55 backdrop-blur-xl">
      <div className="mx-auto flex max-w-7xl items-center justify-between px-6 py-4">
        <div className="flex items-center gap-3">
          <div className="flex h-9 w-9 items-center justify-center rounded-xl bg-white/10 ring-1 ring-white/15">
            <div className="h-4 w-4 rounded bg-blue-500/70" />
          </div>
          <div className="leading-tight">
            <div className="text-base font-semibold tracking-tight">
              {orgLabel} <span className="text-slate-400">- {subtitle}</span>
            </div>
          </div>
        </div>

        <div className="flex items-center gap-3">
          {/* New campaign: ACTIVE only if handler provided */}
          {onNewCampaign ? (
            <button
              type="button"
              onClick={onNewCampaign}
              className="inline-flex items-center gap-2 rounded-xl bg-blue-500/20 px-4 py-2 text-sm font-medium text-blue-100 ring-1 ring-blue-400/25 hover:bg-blue-500/25"
            >
              <PlusIcon className="h-4 w-4" />
              Nouvelle campagne
            </button>
          ) : (
            <div className="inline-flex items-center gap-2 rounded-xl bg-blue-500/10 px-4 py-2 text-sm text-blue-100/80 ring-1 ring-blue-400/15">
              <PlusIcon className="h-4 w-4" />
              Prototype
            </div>
          )}

          <nav className="hidden sm:flex items-center gap-2">
            <a
              href="/curator"
              className={cx(
                "rounded-xl border px-4 py-2 text-sm",
                active === "dashboard"
                  ? "border-white/25 bg-white/[0.10] text-white/90"
                  : "border-white/10 bg-white/[0.06] text-white/85 hover:border-white/20"
              )}
            >
              Dashboard
            </a>

            <a
              href="/curator/exports"
              className={cx(
                "rounded-xl border px-4 py-2 text-sm",
                active === "exports"
                  ? "border-white/25 bg-white/[0.10] text-white/90"
                  : "border-white/10 bg-white/[0.06] text-white/85 hover:border-white/20"
              )}
            >
              Exports
            </a>

            <button
              type="button"
              onClick={onLogout}
              className="rounded-xl border border-white/10 bg-white/[0.06] px-4 py-2 text-sm text-white/85 hover:border-white/20"
            >
              Se déconnecter
            </button>
          </nav>
        </div>
      </div>
    </header>
  );
}
