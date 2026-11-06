import { ConfigModuleOptions } from "@nestjs/config";
import { z } from "zod";

export const EnvConfig: ConfigModuleOptions = {
  isGlobal: true,
  validationSchema: z.object({
    PORT: z.number().default(3000),
  }),
};