"use client"

import { useParams, useRouter } from "next/navigation"
import Image from "next/image"
import { useEffect, useState } from "react"
import Footer from "@/components/global/footer"
import Navbar from "@/components/global/navbar"
import { Button } from "@/components/ui/button"

// Job Detail Data Structure
interface JobDetail {
  id: string
  title: string
  company: string
  logo: string
  location: string
  type: string
  level: string
  salary: string
  description: string
  responsibilities: string[]
  qualifications: string[]
}

// mock data for detail job
const jobData: JobDetail = {
  id: "1",
  title: "Backend Engineer",
  company: "Shopee",
  logo: "/placeholder.svg",
  location: "Jakarta Pusat, Indonesia",
  type: "Full Time",
  level: "Mid-Level",
  salary: "IDR 20.000.000 - 30.000.000/bulan",
  description:
    "Sebagai Backend Engineer di Shopee, Anda akan bertanggung jawab untuk membangun sistem backend yang handal dan scalable untuk mendukung platform e-commerce Shopee.",
  responsibilities: [
    "Merancang dan mengembangkan API untuk layanan backend.",
    "Mengoptimalkan performa sistem dan memastikan skalabilitas.",
    "Menerapkan praktik keamanan terbaik.",
    "Berkolaborasi dengan tim frontend dan mobile.",
  ],
  qualifications: [
    "Minimal S1 di bidang Ilmu Komputer atau terkait.",
    "Pengalaman minimal 2 tahun dalam pengembangan backend.",
    "Menguasai bahasa pemrograman seperti Java, Python, atau Node.js.",
    "Berpengalaman dengan database MySQL, PostgreSQL, atau MongoDB.",
  ],
}

export default function JobDetailPage() {
  const params = useParams()
  const router = useRouter()
  const job = jobData // Gunakan data statis

  return (
    <div className="min-h-screen flex flex-col bg-white">
      {/* Header */}
      <Navbar role="kandidat" />

      {/* Detail Job */}
      <div className="container mx-auto px-4 py-8">
        <div className="p-6">
          <div className="flex items-start gap-4 mb-6 border-b pb-6">
            <div className="bg-red-100 rounded-md p-2">
              <Image src={job.logo} alt={job.company} width={60} height={60} className="rounded-md" />
            </div>
            <div>
              <h1 className="text-3xl font-bold text-[#1e3a5f]">{job.title}</h1>
              <p className="text-2xl pt-2">{job.company}</p>
            </div>
          </div>

          {/* Job Information */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
            <div>
              <p className="text">
                <span className="font-semibold">Lokasi:</span> {job.location}
              </p>
              <p className="text">
                <span className="font-semibold">Tipe Pekerjaan:</span> {job.type}
              </p>
            </div>
            <div>
              <p className="text">
                <span className="font-semibold">Tingkat Posisi:</span> {job.level}
              </p>
              <p className="text">
                <span className="font-semibold">Rentang Gaji:</span> {job.salary}
              </p>
            </div>
          </div>

          {/* Job description */}
          <div className="mb-6">
            <h2 className="text-xl font-semibold text-[#1e3a5f] mb-2">Deskripsi Pekerjaan</h2>
            <p>{job.description}</p>
          </div>

          {/* Job responsibilities */}
          <div className="mb-6">
            <h2 className="text-xl font-semibold text-[#1e3a5f] mb-2">Tanggung Jawab</h2>
            <ul className="list-disc pl-5 space-y-1">
              {job.responsibilities.map((item, index) => (
                <li key={index}>{item}</li>
              ))}
            </ul>
          </div>

          {/* Job qualifications */}
          <div className="mb-6">
            <h2 className="text-xl font-semibold text-[#1e3a5f] mb-2">Kualifikasi</h2>
            <ul className="list-disc pl-5 space-y-1">
              {job.qualifications.map((item, index) => (
                <li key={index}>{item}</li>
              ))}
            </ul>
          </div>

          {/* Button to apply job */}
          <div className="flex pt-10 justify-end">
            <Button className="bg-[#f5efe7] border border-black text-lg text-black font-semibold py-1 px-10 rounded-md hover:bg-[#9db6ca] transition-colors cursor-pointer">
              Lamar
            </Button>
          </div>
        </div>
      </div>

      {/* Footer */}
      <Footer />
    </div>
    
  )
}
