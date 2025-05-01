"use client"

import { useState, useEffect } from "react"
import { useParams, useRouter } from "next/navigation"
import Image from "next/image"
import Link from "next/link"
import { FiArrowLeft } from "react-icons/fi";
import { FaChevronDown, FaChevronCircleUp } from "react-icons/fa";
import { FaChevronUp } from "react-icons/fa6"
import Navbar from "@/components/global/navbar"
import Footer from "@/components/global/footer"


interface Job {
  id: string
  title: string
  company: string
  logo: string
}

interface Applicant {
  id: string
  name: string
  education: string
  photoUrl: string
  score: number
}

const mockApplicant: Applicant[] = [
  {
    id: "applicant-1",
    name: "Aditya Pratama 1",
    education: "Fresh Graduate Teknologi Informasi Universitas Gadjah Mada",
    photoUrl: "/placeholder.svg?height=40&width=40",
    score: 98,
  },
  {
    id: "applicant-2",
    name: "Aditya Pratama 2",
    education: "Fresh Graduate Teknologi Informasi Universitas Gadjah Mada",
    photoUrl: "/placeholder.svg?height=40&width=40",
    score: 94,
  },
  {
    id: "applicant-3",
    name: "Aditya Pratama 3",
    education: "Fresh Graduate Teknologi Informasi Universitas Gadjah Mada",
    photoUrl: "/placeholder.svg?height=40&width=40",
    score: 88,
  },
]

export default function ApplicantListPage() {
  const params = useParams()
  const router = useRouter()
  const { id } = params

  const mockJob: Job = {
    id: id as string,
    title: "Backend Engineer",
    company: "Shopee",
    logo: "/placeholder.svg?height=60&width=60",
  }

  const [job, setJob] = useState<Job>(mockJob);
  const [applicants, setApplicants] = useState<Applicant[]>([])
  const [showJobInfo, setShowJobInfo] = useState(false)

  useEffect(() => {
    setApplicants(mockApplicant);
  }, []);

  return (
    <div className="min-h-screen flex flex-col bg-white">
      {/* Header */}
      <Navbar role="perekrut"/>

      <div className="container mx-auto px-10 py-6 flex-1">
        {/* Back button */}
        <div className="mb-6">
          <button onClick={() => router.back()} className="flex items-center text-[#1e3a5f] hover:underline">
            <FiArrowLeft className="h-4 w-4 mr-1" />
            Kembali
          </button>
        </div>

        {/* Job header */}
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center mb-6">
          <div className="flex items-center gap-4 mb-4 md:mb-0">
            <div className="bg-red-100 rounded-md p-2">
              <Image
                src={job.logo || "/placeholder.svg"}
                alt={job.company}
                width={60}
                height={60}
                className="rounded-md"
              />
            </div>
            <div>
              <h1 className="text-2xl font-bold text-[#1e3a5f]">{job.title}</h1>
              <p className="text-lg">{job.company}</p>
            </div>
          </div>

          <button
            className="flex items-center gap-2 bg-white border border-gray-300 text-gray-700 px-4 py-2 rounded-md hover:bg-gray-50 transition-colors"
          >
            hapus loker
          </button>
        </div>

        {/* Job info dropdown */}
        <div className="mb-6 border border-black rounded-full overflow-hidden">
          <button
            className="w-full flex items-center justify-between bg-[#F5EFE7] px-4 py-3 text-[#1e3a5f] transition-colors"
          >
            <span className="font-medium">informasi loker</span>
            {showJobInfo ? <FaChevronUp className="h-5 w-5" /> : <FaChevronDown className="h-5 w-5" />}
          </button>

          {showJobInfo && (
            <div className="p-4 border-t">
              {/* Konten informasi loker akan ditampilkan di sini */}
              <p>Detail informasi loker akan ditampilkan di sini.</p>
            </div>
          )}
        </div>

        {/* Applicants list */}
        <div>
          {applicants.length === 0 ? (
            <div className="text-center py-8 bg-gray-50 rounded-md">
              <p className="text-gray-500">Belum ada pelamar untuk lowongan ini</p>
            </div>
          ) : (
            <div className="space-y-4">
              {applicants.map((applicant) => (
                <div
                  key={applicant.id}
                  className="bg-white rounded-lg border border-2 border-gray-400 p-4 flex items-center justify-between"
                >
                  <div className="flex items-center gap-4">
                    <div className="relative h-12 w-12 rounded-full overflow-hidden bg-gray-200">
                      <Image
                        src={applicant.photoUrl || "/placeholder.svg"}
                        alt={applicant.name}
                        fill
                        className="object-cover"
                      />
                    </div>
                    <div>
                      <h3 className="font-bold text-lg text-[#1e3a5f]">{applicant.name}</h3>
                      <p className="text-sm text-gray-600">{applicant.education}</p>
                      <Link href={`/perekrut/pelamar/${applicant.id}`} className="text-xs text-[#1e3a5f] hover:underline">
                        lihat profil
                      </Link>
                    </div>
                  </div>
                  <div className="text-2xl font-bold text-[#1e3a5f]">{applicant.score}</div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Footer */}
      <Footer />
    </div>
)}