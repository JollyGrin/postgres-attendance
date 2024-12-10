"use client";

import { IconChevDown } from "@/components/atoms/icons";
import Layout from "@/components/molecules/Layout";
import { OrgPanels } from "@/components/molecules/OrgPanels";
import Link from "next/link";
import { useSearchParams } from "next/navigation";

export default function OrgPage() {
  const [paramOrgName] = useSearchParams().getAll("") ?? [];
  console.log(`org slug: ${paramOrgName}`);
  return <Layout top={<Top />} body={<Body />} />;
}

const Top = () => (
  <>
    <Link href="/">
      <span className="hover:opacity-50 transition-opacity">org</span>
    </Link>
    <span className="text-border font-bold">/</span>
    <span>innkeeper</span>
    <IconChevDown color="var(--color-border)" />
  </>
);

const Body = () => {
  return (
    <>
      <OrgPanels
        org={{
          img: "logos/innkeeper.png",
          name: "innkeeper.eth",
          slug: "innkeeper",
        }}
        locations={[
          { location: "theinn", name: "hyperfy" },
          { location: "137,-2", name: "dcl" },
        ]}
      />
    </>
  );
};
