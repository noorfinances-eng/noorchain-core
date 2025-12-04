
export default function Footer() {
  return (
    <footer className="w-full border-t mt-16 py-6 px-8 text-sm text-gray-600">
      <div className="max-w-4xl mx-auto flex flex-col gap-3">

        <div className="font-semibold">NOORCHAIN — Social Signal Blockchain</div>

        <div className="flex gap-6">
          <a href="/legal" className="underline">Legal Notices</a>
          <a href="/legal" className="underline">Compliance & Risks</a>
          <a href="/legal" className="underline">Privacy</a>
        </div>

        <div className="flex gap-6">
          <a href="https://github.com" target="_blank" className="underline">
            GitHub
          </a>
          <a href="https://x.com" target="_blank" className="underline">
            X (Twitter)
          </a>
          <a href="mailto:contact@noorchain.org" className="underline">
            Contact
          </a>
        </div>
      </div>
    </footer>
  );
}
