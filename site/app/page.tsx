import Image from "next/image";
import Reveal from "../components/ui/Reveal";

export default function HomePage() {
  return (
    <main className="w-full">
      {/* HERO SECTION — premium depth + PoSS-inspired "signal field" (no globe) */}
      <section className="relative isolate w-full overflow-hidden text-white">
        {/* Base gradient */}
        <div className="absolute inset-0 bg-gradient-to-br from-[#1A6AFF] to-[#00D1B2]" />

        {/* Depth: soft radial shading */}
        <div className="absolute inset-0 bg-[radial-gradient(900px_circle_at_18%_18%,rgba(255,255,255,0.18),transparent_55%),radial-gradient(900px_circle_at_85%_25%,rgba(0,0,0,0.18),transparent_55%),radial-gradient(1100px_circle_at_50%_115%,rgba(0,0,0,0.28),transparent_62%)]" />

        {/* Subtle grid */}
        <div className="absolute inset-0 opacity-[0.14] bg-[linear-gradient(to_right,rgba(255,255,255,0.35)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.35)_1px,transparent_1px)] bg-[size:52px_52px]" />

        {/* Signal field (PoSS-inspired): orbs + connections + gentle pulse */}
        <div className="pointer-events-none absolute inset-0">
          {/* Large atmospheric orbs */}
          <div
            className="absolute -top-24 -left-24 h-[420px] w-[420px] rounded-full bg-white/22 blur-3xl noor-signal-orb"
            style={
              {
                ["--x" as any]: "26px",
                ["--y" as any]: "-18px",
                ["--d" as any]: "18s",
              } as React.CSSProperties
            }
          />
          <div
            className="absolute -bottom-28 left-1/3 h-[520px] w-[520px] rounded-full bg-white/16 blur-3xl noor-signal-orb"
            style={
              {
                ["--x" as any]: "18px",
                ["--y" as any]: "20px",
                ["--d" as any]: "26s",
              } as React.CSSProperties
            }
          />
          <div
            className="absolute top-10 -right-28 h-[460px] w-[460px] rounded-full bg-white/14 blur-3xl noor-signal-orb"
            style={
              {
                ["--x" as any]: "-22px",
                ["--y" as any]: "14px",
                ["--d" as any]: "22s",
              } as React.CSSProperties
            }
          />

          {/* Smaller "signal nodes" */}
          <div
            className="absolute top-24 left-[12%] h-24 w-24 rounded-full bg-white/16 blur-2xl noor-signal-orb"
            style={
              {
                ["--x" as any]: "14px",
                ["--y" as any]: "10px",
                ["--d" as any]: "16s",
              } as React.CSSProperties
            }
          />
          <div
            className="absolute top-44 right-[18%] h-28 w-28 rounded-full bg-white/14 blur-2xl noor-signal-orb"
            style={
              {
                ["--x" as any]: "-12px",
                ["--y" as any]: "-10px",
                ["--d" as any]: "19s",
              } as React.CSSProperties
            }
          />
          <div
            className="absolute bottom-24 left-[22%] h-24 w-24 rounded-full bg-white/12 blur-2xl noor-signal-orb"
            style={
              {
                ["--x" as any]: "10px",
                ["--y" as any]: "-12px",
                ["--d" as any]: "17s",
              } as React.CSSProperties
            }
          />

          {/* Connections (subtle) */}
          <div
            className="noor-connection left-24 top-44 w-[420px]"
            style={{ transform: "rotate(8deg)" }}
          />
          <div
            className="noor-connection right-24 top-56 w-[360px]"
            style={{ transform: "rotate(-10deg)" }}
          />
          <div
            className="noor-connection left-40 bottom-40 w-[520px]"
            style={{ transform: "rotate(-6deg)" }}
          />

          {/* Signal pulse */}
          <div className="noor-pulse left-24 top-52" />
        </div>

        {/* Content */}
        <div className="relative">
          <div className="container py-14 sm:py-16 md:py-24">
            <div className="max-w-3xl">
              <div className="flex flex-col sm:flex-row items-start sm:items-center gap-3 sm:gap-4 mb-5">
                <div className="relative">
                  <div className="absolute -inset-3 rounded-2xl bg-white/15 blur-xl" />
                  <div className="relative rounded-2xl bg-white/10 ring-1 ring-white/25 p-2">
                    <Image
                      src="/logo-white.svg"
                      alt="NOORCHAIN Logo"
                      width={52}
                      height={52}
                      priority
                    />
                  </div>
                </div>

                <h1 className="text-3xl sm:text-4xl md:text-6xl font-extrabold tracking-tight">
                  NOORCHAIN
                </h1>
              </div>

              <p className="text-base sm:text-lg md:text-xl text-white/90 mb-3 font-medium">
                A Human-Centered Blockchain for Social Signals
              </p>

              <p className="text-xs sm:text-sm text-white/80 mb-6 font-medium">
                Private mainnet-like environment — controlled operation, non-public
                by design.
              </p>

              <p className="text-sm sm:text-base md:text-lg text-white/90 leading-relaxed mb-8">
                A Social Signal Blockchain powered by PoSS. Designed for transparent
                participation, curator validation, and a fixed-supply digital model
                free from financial speculation.
              </p>

              <div className="flex flex-col sm:flex-row flex-wrap gap-3 sm:gap-4">
                <a
                  href="/technology"
                  className="w-full sm:w-auto text-center px-6 py-3 rounded-md text-sm md:text-base font-medium
                             bg-white text-[#0B1B3A] ring-1 ring-white/30
                             shadow-[0_12px_34px_rgba(0,0,0,0.18)]
                             transition-all duration-200 hover:-translate-y-0.5 hover:shadow-[0_16px_46px_rgba(0,0,0,0.22)]"
                >
                  Explore Technology
                </a>

                <a
                  href="/genesis"
                  className="w-full sm:w-auto text-center px-6 py-3 rounded-md text-sm md:text-base font-medium
                             border border-white/70 text-white
                             bg-white/0 transition-all duration-200
                             hover:bg-white/10 hover:-translate-y-0.5"
                >
                  Genesis Overview
                </a>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* STATUS + INTRO + PoSS FRAMING — 1 full-width + 2 half-width */}
      <section className="container py-10 md:py-14">
        <div className="grid grid-cols-1 gap-6">
          {/* ROW 1: FULL WIDTH */}
          <Reveal delayMs={0}>
            <div
              className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm
                         transition-all duration-200 hover:shadow-md hover:-translate-y-0.5"
            >
              <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-4">
                <h2 className="text-lg sm:text-xl font-bold text-gray-900">
                  Current Project Status
                </h2>
                <p className="text-xs sm:text-sm text-gray-500">
                  Last updated: 2026-01-01
                </p>
              </div>

              <ul className="text-sm sm:text-base text-gray-700 leading-relaxed space-y-2">
                <li>
                  <span className="font-semibold text-gray-900">Network:</span>{" "}
                  Private mainnet-like environment (continuous operation)
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Consensus:</span>{" "}
                  Permissioned BFT
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Layer 1:</span>{" "}
                  Sovereign EVM L1
                </li>
                <li>
                  <span className="font-semibold text-gray-900">PoSS:</span>{" "}
                  Application layer for governance and verifiable contribution
                  signals (not consensus)
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Public access:</span>{" "}
                  Limited by design until feature completeness and security review
                </li>
                <li>
                  <span className="font-semibold text-gray-900">Reference build:</span>{" "}
                  M10-MAINNETLIKE-STABLE / M11-DAPPS-STABLE / M12.2-WORLDSTATE-RPC-NONCE-BALANCE
                </li>
              </ul>
            </div>
          </Reveal>

          {/* ROW 2: TWO COLUMNS (md+) */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <Reveal delayMs={80}>
              <div
                className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm
                           transition-all duration-200 hover:shadow-md hover:-translate-y-0.5"
              >
                <h2 className="text-lg sm:text-xl font-bold text-gray-900 mb-4">
                  A New Approach to Blockchain Design
                </h2>

                <p className="text-sm sm:text-base text-gray-700 leading-relaxed">
                  NOORCHAIN introduces a mission-driven blockchain architecture
                  focused on verified social participation rather than financial
                  speculation. Powered by the PoSS protocol and aligned with Legal
                  Light CH, it provides a transparent and sustainable digital
                  infrastructure for curators, participants, institutions and
                  communities.
                </p>
              </div>
            </Reveal>

            <Reveal delayMs={140}>
              <div
                className="rounded-xl border border-gray-200 bg-white p-6 md:p-8 shadow-sm
                           transition-all duration-200 hover:shadow-md hover:-translate-y-0.5"
              >
                <div className="flex flex-col sm:flex-row sm:items-baseline sm:justify-between gap-2 mb-4">
                  <h2 className="text-lg sm:text-xl font-bold text-gray-900">
                    PoSS Framing (Non-Consensus)
                  </h2>
                  <a
                    href="/poss"
                    className="text-sm font-medium text-blue-700 hover:text-blue-900 transition"
                  >
                    Read PoSS framing
                  </a>
                </div>

                <ul className="text-sm sm:text-base text-gray-700 leading-relaxed space-y-2">
                  <li>
                    <span className="font-semibold text-gray-900">PoSS is not consensus.</span>{" "}
                    Network security is provided by a permissioned BFT consensus layer.
                  </li>
                  <li>
                    <span className="font-semibold text-gray-900">PoSS is an application layer.</span>{" "}
                    It structures governance, coordination, and verifiable contribution signals.
                  </li>
                  <li>
                    <span className="font-semibold text-gray-900">Economic posture:</span>{" "}
                    NUR is the native gas token; NOORCHAIN does not offer returns or custody.
                  </li>
                </ul>
              </div>
            </Reveal>
          </div>
        </div>
      </section>
    </main>
  );
}
