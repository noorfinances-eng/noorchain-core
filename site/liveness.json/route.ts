export const runtime = "nodejs";

export async function GET() {
  // IMPORTANT: do not expose internal RPC/P2P details.
  // Only proxy the minimal public signal.
  const upstream = process.env.LIVENESS_UPSTREAM_URL;

  if (!upstream) {
    return new Response(JSON.stringify({ error: "liveness_upstream_not_set" }), {
      status: 503,
      headers: {
        "content-type": "application/json; charset=utf-8",
        "cache-control": "no-store",
      },
    });
  }

  try {
    const r = await fetch(upstream, { cache: "no-store" });
    const txt = await r.text();

    // Pass-through (expected JSON); keep it minimal.
    return new Response(txt, {
      status: r.status,
      headers: {
        "content-type": "application/json; charset=utf-8",
        "cache-control": "no-store",
      },
    });
  } catch {
    return new Response(JSON.stringify({ error: "liveness_unreachable" }), {
      status: 502,
      headers: {
        "content-type": "application/json; charset=utf-8",
        "cache-control": "no-store",
      },
    });
  }
}
