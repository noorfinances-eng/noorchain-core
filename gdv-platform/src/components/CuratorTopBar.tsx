"use client";

import React from "react";

type Active = "dashboard" | "exports";

function cx(...parts: Array<string | false | null | undefined>) {
  return parts.filter(Boolean).join(" ");
}

type IconName = "plus" | "clock" | "user" | "menu" | "cloud" | "gear";

function Icon({ name, className }: { name: IconName; className?: string }) {
  const base = cx("h-5 w-5", className);
  switch (name) {
    case "plus":
      return (
        <svg className={base} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="M12 5v14M5 12h14" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
        </svg>
      );
    case "clock":
      return (
        <svg className={base} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path
            d="M12 8v5l3 2"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"
            stroke="currentColor"
            strokeWidth="2"
          />
        </svg>
      );
    case "user":
      return (
        <svg className={base} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path
            d="M20 21a8 8 0 0 0-16 0"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
          />
          <path
            d="M12 13a5 5 0 1 0-5-5a5 5 0 0 0 5 5Z"
            stroke="currentColor"
            strokeWidth="2"
          />
        </svg>
      );
    case "menu":
      return (
        <svg className={base} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path d="M4 7h16M4 12h16M4 17h16" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
        </svg>
      );
    case "cloud":
      return (
        <svg className={base} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path
            d="M7 18h10a4 4 0 0 0 .8-7.92A6 6 0 0 0 6.2 9.6A3.5 3.5 0 0 0 7 18Z"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      );
    case "gear":
      return (
        <svg className={base} viewBox="0 0 24 24" fill="none" aria-hidden="true">
          <path
            d="M12 15.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Z"
            stroke="currentColor"
            strokeWidth="2"
          />
          <path
            d="M19.4 15a7.8 7.8 0 0 0 .1-1l2-1.2l-2-3.4l-2.3.6a7.9 7.9 0 0 0-1.7-1l-.3-2.3H10.8l-.3 2.3a7.9 7.9 0 0 0-1.7 1l-2.3-.6l-2 3.4l2 1.2a7.8 7.8 0 0 0 .1 1a7.8 7.8 0 0 0-.1 1l-2 1.2l2 3.4l2.3-.6a7.9 7.9 0 0 0 1.7 1l.3 2.3h4.4l.3-2.3a7.9 7.9 0 0 0 1.7-1l2.3.6l2-3.4l-2-1.2a7.8 7.8 0 0 0-.1-1Z"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      );
  }
}

export function CuratorTopBar({
  active,
  onLogout,
  subtitle = "Graine de Vie",
  orgLabel = "Curators Hub",
  showPrototypeControls = false,
}: {
  active: Active;
  onLogout: () => void;
  subtitle?: string;
  orgLabel?: string;
  showPrototypeControls?: boolean;
}) {
  const disabled = "opacity-60 cursor-not-allowed";

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
          {showPrototypeControls ? (
            <button
              type="button"
              className={cx(
                "inline-flex items-center gap-2 rounded-xl bg-blue-500/20 px-4 py-2 text-sm font-medium text-blue-100 ring-1 ring-blue-400/25 hover:bg-blue-500/25"
              )}
            >
              <Icon name="plus" className="h-4 w-4" />
              Nouvelle campagne
            </button>
          ) : (
            <div className={cx("inline-flex items-center gap-2 rounded-xl bg-blue-500/10 px-4 py-2 text-sm text-blue-100/80 ring-1 ring-blue-400/15", disabled)}>
              <Icon name="plus" className="h-4 w-4" />
              Nouvelle campagne
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

          <div className="flex items-center gap-2 text-slate-300">
            <div className={cx("rounded-xl bg-white/5 p-2 ring-1 ring-white/10", disabled)} aria-label="Historique">
              <Icon name="clock" />
            </div>
            <div className={cx("rounded-xl bg-white/5 p-2 ring-1 ring-white/10", disabled)} aria-label="Compte">
              <Icon name="user" />
            </div>
            <div className={cx("rounded-xl bg-white/5 p-2 ring-1 ring-white/10", disabled)} aria-label="Menu">
              <Icon name="menu" />
            </div>
            <div className={cx("rounded-xl bg-white/5 p-2 ring-1 ring-white/10", disabled)} aria-label="Sync">
              <Icon name="cloud" />
            </div>
            <div className={cx("rounded-xl bg-white/5 p-2 ring-1 ring-white/10", disabled)} aria-label="Paramètres">
              <Icon name="gear" />
            </div>
          </div>
        </div>
      </div>
    </header>
  );
}
