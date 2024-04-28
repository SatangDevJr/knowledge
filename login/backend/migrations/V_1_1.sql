USE [exam2022]
GO
/****** Object:  Table [dbo].[TB_MAS_User]    Script Date: 15/12/2022 4:33:54 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TB_MAS_User](
	[ID] [int] IDENTITY(1,1) NOT NULL,
	[UserName] [nvarchar](50) NOT NULL,
	[Password] [nvarchar](250) NOT NULL,
	[CreatedDate] [datetime] NULL,
	[CreatedByUserId] [int] NULL,
	[UpdatedDate] [datetime] NULL,
	[UpdatedByUserId] [int] NULL,
	[DelFlag] [bit] NULL,
 CONSTRAINT [PK_TB_MAS_User] PRIMARY KEY CLUSTERED 
(
	[ID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
