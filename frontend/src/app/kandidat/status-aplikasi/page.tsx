"use client"

import { useState, useEffect } from "react"
import Image from "next/image"
import Link from "next/link"
import Navbar from "@/components/global/navbar"
import Footer from "@/components/global/footer"

// Mock data for application status
const mockApplications = [
  {
    id: "app-1",
    jobTitle: "Backend Engineer",
    company: "Shopee",
    logo: "/placeholder.svg?height=40&width=40",
    status: "proses review",
    appliedDate: "2023-03-15",
  },
  {
    id: "app-2",
    jobTitle: "Frontend Developer",
    company: "Tokopedia",
    logo: "/placeholder.svg?height=40&width=40",
    status: "ditolak",
    appliedDate: "2023-02-20",
  },
  {
    id: "app-3",
    jobTitle: "Data Scientist",
    company: "Gojek",
    logo: "/placeholder.svg?height=40&width=40",
    status: "interview",
    appliedDate: "2023-01-10",
  },
  {
    id: "app-4",
    jobTitle: "Mobile Developer",
    company: "Bukalapak",
    logo: "/placeholder.svg?height=40&width=40",
    status: "diterima",
    appliedDate: "2023-03-01",
  },
]

export default function ApplicationStatusPage() {
  const [applications, setApplications] = useState(mockApplications) // Use mock data
  const [isLoading, setIsLoading] = useState(false) // Directly set loading to false as no fetching occurs

  const getStatusBadgeClass = (status: string) => {
    switch (status) {
      case "proses review":
        return "bg-blue-100 text-blue-800"
      case "ditolak":
        return "bg-red-100 text-red-800"
      case "diterima":
        return "bg-green-100 text-green-800"
      case "interview":
        return "bg-yellow-100 text-yellow-800"
      default:
        return "bg-gray-100 text-gray-800"
    }
  }

  if (isLoading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <h1 className="text-2xl font-bold text-center text-[#1e3a5f] mb-8">Status Aplikasi</h1>
        <div className="animate-pulse space-y-4">
          {[1, 2, 3].map((i) => (
            <div key={i} className="bg-gray-100 p-4 rounded-lg">
              <div className="h-6 bg-gray-200 rounded w-1/3 mb-2"></div>
              <div className="h-4 bg-gray-200 rounded w-1/4"></div>
            </div>
          ))}
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen flex flex-col bg-white">
			{/* Header */}
      <Navbar role="kandidat"/>

			{/* Application status */}
      <div className="container mx-auto px-4 py-8">
        <h1 className="text-4xl font-bold text-center text-[#1e3a5f] mb-8">Status Aplikasi</h1>
          <div className="max-w-4xl mx-auto">
            <div className="border-b border-gray-300 mb-6"></div>
            {applications.length === 0 ? (
            <p className="text-center text-gray-500 py-8">
                Anda belum mengajukan lamaran pekerjaan.
                <br />
                <Link href="/kandidat" className="text-[#1e3a5f] hover:underline mt-2 inline-block">
                Cari pekerjaan sekarang
                </Link>
            </p>
            ) : (
            <div className="space-y-4">
                {applications.map((application) => (
                <div
                    key={application.id}
                    className="bg-[#f5efe7]/40 border border-black rounded-lg p-6 flex items-center justify-between"
                >
                    <div className="flex items-center">
                    <div className="bg-red-100 rounded-md p-2 mr-4">
                        <Image
                        src={application.logo || "/placeholder.svg"}
                        alt={application.company}
                        width={40}
                        height={40}
                        className="rounded-md"
                        />
                    </div>
                    <div>
                        <h3 className="font-bold text-lg text-[#1e3a5f]">{application.jobTitle}</h3>
                        <p className="text-black">{application.company}</p>
                    </div>
                    </div>
                    <div>
                    <span
                        className={`px-3 py-1 rounded-full text-sm border font-medium ${getStatusBadgeClass(application.status)}`}
                    >
                        {application.status}
                    </span>
                    </div>
                </div>
                ))}
            </div>
            )}
        </div>
      </div>

      {/* Footer */}
      <Footer />
    </div>
  )
}
