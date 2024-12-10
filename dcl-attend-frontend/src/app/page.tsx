"use client";

import Layout from "@/components/molecules/Layout";
import Link from "next/link";

export default function Home() {
  return <Layout top={<Top />} body={<Body />} />;
}

const Top = () => (
  <>
    <span>orgs</span>
  </>
);

const Body = () => (
  <>
    <Link href="/org?=innkeeper">
      <div className="bg-panel p-4 border-border border-2 rounded-lg flex items-center gap-4 text-xl hover:opacity-50 transition-opacity">
        <img src="logos/innkeeper.png" height="40px" width="40px" />
        <p>innkeeper.eth</p>
      </div>
    </Link>
  </>
);
