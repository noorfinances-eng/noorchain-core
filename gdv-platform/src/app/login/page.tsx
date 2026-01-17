"use client";

import { useMemo, useState } from "react";
import { useRouter } from "next/navigation";

type Role = "curator" | "dev";

export default function LoginPage() {
  const router = useRouter();
  const users = useMemo(
    () => ({
      curator: { username: "grainedevie", password: "noorchain2!", landing: "/curator" },
      dev: { username: "developper", password: "noorchain22!", landing: "/dev" },
    }),
    []
  );

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  function setRole(role: Role) {
    // For the next step (middleware RBAC), we set both cookie + localStorage (dev-only prototype).
    document.cookie = `gdv_role=${role}; path=/; samesite=lax`;
    localStorage.setItem("gdv_role", role);
  }

  function matchRole(): Role | null {
    if (username === users.curator.username && password === users.curator.password) return "curator";
    if (username === users.dev.username && password === users.dev.password) return "dev";
    return null;
  }


  function sanitizeNext(raw: string | null): string | null {
    if (!raw) return null;
    // Only allow internal app routes we actually serve (no open redirect).
    if (raw.startsWith("/curator") || raw.startsWith("/dev")) return raw;
    return null;
  }

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError(null);
    setIsSubmitting(true);

    try {
      const role = matchRole();
      if (!role) {
        setError("Identifiants invalides.");
        return;
      }
      setRole(role);
      const params = new URLSearchParams(window.location.search);
        const next = sanitizeNext(params.get("next"));
        const fallback = role === "dev" ? users.dev.landing : users.curator.landing;
        router.push(next ?? fallback);
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <main className="min-h-screen bg-[#071025] text-white">
      <div className="mx-auto flex min-h-screen max-w-lg items-center px-6">
        <div className="w-full rounded-2xl border border-white/10 bg-white/[0.06] p-6 shadow-[0_8px_30px_rgba(0,0,0,0.35)] backdrop-blur-xl">
          <div className="mb-6">
            <div className="text-xs tracking-widest text-white/60">GDV · Curators Hub</div>
            <h1 className="mt-2 text-2xl font-semibold">Connexion</h1>
            <p className="mt-2 text-sm text-white/70">
              Prototype local (dev-only). Accès Curator ou Dev.
            </p>
          </div>

          <form onSubmit={onSubmit} className="space-y-4">
            <div>
              <label className="text-sm text-white/80">Nom d’utilisateur</label>
              <input
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="mt-2 w-full rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm outline-none focus:border-white/25"
                placeholder="Nom d’utilisateur"
                autoComplete="username"
              />
            </div>

            <div>
              <label className="text-sm text-white/80">Mot de passe</label>
              <input
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                type="password"
                className="mt-2 w-full rounded-xl border border-white/10 bg-white/[0.06] px-4 py-3 text-sm outline-none focus:border-white/25"
                placeholder="Mot de passe"
                autoComplete="current-password"
              />
            </div>

            {error ? (
              <div className="rounded-xl border border-red-500/30 bg-red-500/10 px-4 py-3 text-sm text-red-100">
                {error}
              </div>
            ) : null}

            <button
              type="submit"
              disabled={isSubmitting}
              className="w-full rounded-xl bg-white text-[#071025] px-4 py-3 text-sm font-semibold disabled:opacity-60"
            >
              {isSubmitting ? "Connexion..." : "Se connecter"}
            </button>

            <div className="pt-2 text-xs text-white/50">
              RBAC (blocage des routes) sera appliqué à l’étape suivante via middleware.
            </div>
          </form>
        </div>
      </div>
    </main>
  );
}
