import { ConfigModuleOptions, registerAs } from "@nestjs/config";
import { z } from "zod";

const db = registerAs('db', () => ({
  url: process.env.DB_URL,
}));

export const EnvConfig: ConfigModuleOptions = {
  isGlobal: true,
  validationSchema: z.object({
    PORT: z.coerce.number().default(3000),
    DB_URL: z.string().min(1),
  }),
  load: [db],
};